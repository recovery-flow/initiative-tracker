package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/internal/service/requests"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParticipantCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.ParticipantCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	participantId, err := uuid.Parse(req.Data.Attributes.UserId)
	if err != nil {
		log.WithError(err).Error("Failed to parse participant id")
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

	participantRole, err := roles.StringToRoleTeam(req.Data.Attributes.Role)
	if err != nil {
		log.WithError(err).Error("Failed to parse role")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if roles.CompareRolesTeam(initiator.Role, participantRole) < 0 || participantRole == roles.RoleTeamOwner {
		log.Error("User has no rights to create participant")
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to create participant"))
		return
	}

	res, err := server.MongoDB.Participants.New().Insert(r.Context(), models.Participant{
		InitiativeID: iniId,
		UserID:       participantId,
		Role:         participantRole,
		FirstName:    req.Data.Attributes.FirstName,
		SecondName:   req.Data.Attributes.SecondName,
		ThirdName:    req.Data.Attributes.ThirdName,
		DisplayName:  req.Data.Attributes.DisplayName,
		Position:     req.Data.Attributes.Position,
		Desc:         req.Data.Attributes.Desc,
		Verified:     false,

		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			httpkit.RenderErr(w, problems.Conflict("User already exists in initiative"))
			return
		}
		log.WithError(err).Error("Failed to create participant")
		httpkit.RenderErr(w, problems.InternalError("Failed to create participant"))
		return
	}

	httpkit.Render(w, responses.Participant(*res))
}
