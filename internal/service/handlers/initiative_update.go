package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/requests"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitiativeUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewInitiativeUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if chi.URLParam(r, "initiative_id") != req.Data.Id {
		log.Error("Request id does not match URL id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"id": validation.NewError("id", "Request id does not match URL id"),
		})...)
		return

	}

	iniId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"initiative_id": validation.NewError("initiative_id", "Invalid initiative id"),
		})...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	filters := make(map[string]any)
	filters["initiative_id"] = iniId
	filters["user_id"] = initiatorId

	participant, err := server.MongoDB.Participants.New().Filter(filters).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get initiative")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if roles.CompareRolesTeam(participant.Role, roles.RoleTeamModer) <= 0 {
		log.Error("User has no rights to update initiative")
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to update initiative"))
		return
	}

	stmt := make(map[string]any)
	if req.Data.Attributes.Name != nil {
		stmt["name"] = req.Data.Attributes.Name
	}
	if req.Data.Attributes.Desc != nil {
		stmt["desc"] = req.Data.Attributes.Desc
	}
	if req.Data.Attributes.Goal != nil {
		stmt["goal"] = req.Data.Attributes.Goal
	}
	if req.Data.Attributes.Location != nil {
		stmt["location"] = req.Data.Attributes.Location
	}
	if req.Data.Attributes.Status != nil {
		stmt["status"] = req.Data.Attributes.Status
	}

	res, err := server.MongoDB.Initiative.New().Filter(map[string]any{
		"_id": iniId,
	}).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to update initiative")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	log.Infof("Initiative %s updated %s", iniId, initiatorId)
	httpkit.Render(w, responses.Initiative(*res))
}
