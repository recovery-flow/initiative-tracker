package service

import (
	"context"

	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/handlers"
	"github.com/recovery-flow/roles"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, service))
	authMW := service.TokenManager.AuthMdl(service.Config.JWT.AccessToken.SecretKey)
	adminGrant := service.TokenManager.RoleGrant(service.Config.JWT.AccessToken.SecretKey, string(roles.RoleUserAdmin))

	r.Route("/initiative-tracker", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/private", func(r chi.Router) {
				r.Use(authMW)

				r.Route("/initiative", func(r chi.Router) {
					r.Post("/", handlers.InitiativeCreate)

					r.Route("/{initiative_id}", func(r chi.Router) {
						r.Put("/", handlers.InitiativeUpdate)
						r.Patch("/wallets", handlers.InitiativeWalletsUpdate)
						r.Patch("/organizations", handlers.InitiativeUpdateOrgsMembers)
					})
				})
			})

			r.Route("/public", func(r chi.Router) {
				r.Route("/initiative", func(r chi.Router) {
					r.Get("/", nil) // search initiatives by filters

					r.Route("/{initiative_id}", func(r chi.Router) {
						r.Get("/", handlers.InitiativeGet)
					})
				})
			})

			r.Route("/admin", func(r chi.Router) {
				r.Use(adminGrant)

				r.Route("/initiative", func(r chi.Router) {
					r.Route("/{initiative_id}", func(r chi.Router) {
						r.Patch("/{value}", nil) //admin update initiative
					})
				})
			})
		})
	})
	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}
