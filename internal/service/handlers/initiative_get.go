package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitiativeGet(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	iniID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"initiative_id": validation.NewError("initiative_id", "Invalid initiative id"),
		})...)
		return
	}

	filters := make(map[string]any)
	filters["_id"] = iniID

	initiative, err := server.MongoDB.Initiative.New().Filter(filters).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get initiative")
		httpkit.RenderErr(w, problems.InternalError("Failed to get initiative"))
		return
	}

	httpkit.Render(w, responses.Initiative(*initiative))
}
