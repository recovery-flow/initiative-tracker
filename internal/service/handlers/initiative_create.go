package handlers

import (
	"net/http"
	"strings"
	"time"

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
)

func InitiativeCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewInitiativeCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	owner := req.Data.Attributes.Owner

	initiative := models.Initiative{
		ID:       primitive.NewObjectID(),
		Name:     req.Data.Attributes.Name,
		Desc:     req.Data.Attributes.Desc,
		Goal:     req.Data.Attributes.Goal,
		Verified: false,
		Location: req.Data.Attributes.Location,
		Status:   models.StatusInactive,
		Tags:     []models.Tag{},
		ChatID:   primitive.NilObjectID,

		Likes:   0,
		Reposts: 0,
		Reports: 0,

		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	res, err := server.MongoDB.Initiative.New().Insert(r.Context(), initiative)
	if err != nil {
		if strings.HasPrefix(err.Error(), "failed to insert initiative") {
			httpkit.RenderErr(w, problems.InternalError("Failed to insert initiative"))
			return
		}
		log.Infof("Failed to create initiative: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ownerPrt := models.Participant{
		UserID:       userId,
		InitiativeID: res.ID,
		FirstName:    owner.FirstName,
		SecondName:   owner.SecondName,
		ThirdName:    owner.ThirdName,
		DisplayName:  owner.DisplayName,
		Position:     owner.Position,
		Desc:         owner.Desc,
		Verified:     false,
		Role:         roles.RoleTeamOwner,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	_, err = server.MongoDB.Participants.New().Insert(r.Context(), ownerPrt)
	if err != nil {
		if strings.HasPrefix(err.Error(), "failed to insert participant") {
			httpkit.RenderErr(w, problems.InternalError("Failed to insert participant"))
			return
		}
		log.Infof("Failed to create participant: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	log.Infof("Initiative created: %v by user %s", res.ID, userId)
	httpkit.Render(w, responses.Initiative(*res))
}
