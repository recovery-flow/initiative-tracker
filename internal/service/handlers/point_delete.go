package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PointDelete(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	iniId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"initiative_id": validation.NewError("initiative_id", "Failed to parse initiative id"),
		})...)
		return
	}

	pointId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "point_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse point id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"point_id": validation.NewError("point_id", "Failed to parse point id"),
		})...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	initiator, err := server.MongoDB.Participants.New().Filter(map[string]any{
		"initiative_id": iniId,
		"user_id":       userId,
	}).Get(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Participant not found %s", userId)
			httpkit.RenderErr(w, problems.NotFound("Participant not found"))
			return
		}
		log.WithError(err).Error("Failed to get point")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if roles.CompareRolesTeam(initiator.Role, roles.RoleTeamAdmin) < 0 {
		log.Errorf("User has no rights to update point %s || %s", initiator.UserID, initiator.Role)
		httpkit.RenderErr(w, problems.Forbidden("User has no rights to update point"))
		return
	}

	curPoint, err := server.MongoDB.Points.New().Filter(map[string]any{
		"_id":           pointId,
		"initiative_id": iniId,
	}).Get(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Point not found %s", pointId)
			httpkit.RenderErr(w, problems.NotFound("Point not found"))
			return
		}
		log.WithError(err).Error("Failed to get point")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if curPoint.Level == 0 {
		log.Errorf("Cannot change level of root point %s", pointId)
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"level": validation.NewError("level", "Cannot change level of root point"),
		})...)
		return
	}

	children, err := server.MongoDB.Points.New().Filter(map[string]any{
		"parent_id":     curPoint.ID,
		"initiative_id": iniId,
	}).Select(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Parent point not found %s", curPoint.ParentID)
			httpkit.RenderErr(w, problems.NotFound("Parent point not found"))
			return
		}
		log.WithError(err).Error("Failed to get parent point")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if len(children) != 0 {
		log.Infof("Cannot delete point with children %v", children)
		log.Errorf("Cannot delete point with children %s", pointId)
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"level": validation.NewError("level", "Cannot delete point with children"),
		})...)
		return
	}

	err = server.MongoDB.Points.New().Filter(map[string]any{
		"_id":           pointId,
		"initiative_id": iniId,
	}).DeleteOne(r.Context())
	if err != nil {
		log.WithError(err).Error("Failed to update point")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusOK)
}
