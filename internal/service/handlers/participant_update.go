package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

func ParticipantUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.ParticipantUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	participantUserId, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse participant user id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	iniId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	initiator, err := server.MongoDB.Participants.New().Filter(map[string]any{
		"initiative_id": iniId,
		"user_id":       initiatorId,
	}).Get(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to get participant")
		httpkit.RenderErr(w, problems.InternalError("Failed to get participant"))
		return
	}

	participant, err := server.MongoDB.Participants.New().Filter(map[string]any{
		"user_id":       participantUserId,
		"initiative_id": iniId,
	}).Get(r.Context())

	var newRole roles.TeamRole
	if req.Data.Attributes.Role != nil {
		newRole, err = roles.StringToRoleTeam(*req.Data.Attributes.Role)
		if err != nil {
			log.WithError(err).Error("Failed to parse role")
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		if roles.CompareRolesTeam(initiator.Role, participant.Role) != 1 || roles.CompareRolesTeam(initiator.Role, newRole) < 0 || newRole == roles.RoleTeamOwner {
			log.Errorf("User has no rights to update participant %s -> %s", initiator.Role, newRole)
			httpkit.RenderErr(w, problems.Forbidden("User has no rights to update participant"))
			return
		}
	}

	stmt := make(map[string]any)
	if req.Data.Attributes.FirstName != nil {
		stmt["first_name"] = req.Data.Attributes.FirstName
		stmt["verified"] = false
	}
	if req.Data.Attributes.SecondName != nil {
		stmt["second_name"] = req.Data.Attributes.SecondName
		stmt["verified"] = false
	}
	if req.Data.Attributes.ThirdName != nil {
		stmt["third_name"] = req.Data.Attributes.ThirdName
		stmt["verified"] = false
	}
	if req.Data.Attributes.DisplayName != nil {
		stmt["display_name"] = req.Data.Attributes.DisplayName
	}
	if req.Data.Attributes.Desc != nil {
		stmt["desc"] = req.Data.Attributes.Desc
	}
	if req.Data.Attributes.Position != nil {
		stmt["position"] = req.Data.Attributes.Position
	}
	if req.Data.Attributes.Role != nil {
		stmt["role"] = newRole
	}

	res, err := server.MongoDB.Participants.New().Filter(map[string]any{
		"user_id":       participantUserId,
		"initiative_id": iniId,
	}).UpdateOne(r.Context(), stmt)
	if err != nil {
		log.WithError(err).Error("Failed to create participant")
		httpkit.RenderErr(w, problems.InternalError("Failed to create participant"))
		return
	}

	httpkit.Render(w, responses.Participant(*res))
}
