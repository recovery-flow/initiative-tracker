package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParticipantByOrganization(w http.ResponseWriter, r *http.Request) {
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
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	userID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	res, err := server.MongoDB.Initiative.New().Filter(map[string]any{
		"_id": iniID,
	}).Participants().Filter(map[string]any{
		"user_id": userID,
	}).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get participant")
		httpkit.RenderErr(w, problems.InternalError("Failed to get participant"))
		return
	}

	httpkit.Render(w, responses.Participant(*res))
}
