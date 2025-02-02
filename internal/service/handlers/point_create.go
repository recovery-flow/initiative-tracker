package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func PointCreate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	req, err := requests.NewPointCreate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	iniId, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"initiative_id": validation.NewError("initiative_id", "Failed to parse initiative id"),
		})...)
		return
	}

	userId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		server.Logger.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	var parentID primitive.ObjectID
	if req.Data.Attributes.ParentId != nil {
		parentID, err = primitive.ObjectIDFromHex(*req.Data.Attributes.ParentId)
		if err != nil {
			log.WithError(err).Error("Failed to parse parent id")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"parent_id": validation.NewError("parent_id", "Failed to parse parent id"),
			})...)
			return
		}

		parent, err := server.MongoDB.Points.New().FilterExact(map[string]any{
			"_id": parentID,
		}).Get(r.Context())
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				log.Errorf("Parent point not found %s", parentID)
				httpkit.RenderErr(w, problems.NotFound("Parent point not found"))
				return
			}
			log.WithError(err).Error("Failed to get parent point")
			httpkit.RenderErr(w, problems.InternalError())
			return
		}

		if parent.InitiativeID != iniId {
			err = fmt.Errorf("parent point does not belong to the same initiative")
			log.WithError(err).Error("Parent point does not belong to the same initiative")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"parent_id": validation.NewError("parent_id", "Parent point does not belong to the same initiative"),
			})...)
			return
		}
		if parent.Level >= int(req.Data.Attributes.Level) {
			err = fmt.Errorf("new point level must be greater than parent point level")
			log.WithError(err).Error("New point level must be greater than parent point level")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"level": validation.NewError("level", "New point level must be greater than parent point level"),
			})...)
			return
		}
	}

	if req.Data.Attributes.Level > 0 {
		if parentID.IsZero() {
			err = fmt.Errorf("parent point id is required for non-root points")
			log.WithError(err).Error("Parent point id is required for non-root points")
			httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
				"parent_id": validation.NewError("parent_id", "Parent point id is required for non-root points"),
			})...)
			return
		}
	}

	if req.Data.Attributes.Level < 0 {
		err = fmt.Errorf("point level must be a non-negative integer")
		log.WithError(err).Error("Point level must be a non-negative integer")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"level": validation.NewError("level", "Point level must be a non-negative integer"),
		})...)
		return
	}

	initiator, err := server.MongoDB.Participants.New().FilterExact(map[string]any{
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

	zero := 0.0
	point := models.Point{
		ID:             primitive.NewObjectID(),
		InitiativeID:   iniId,
		ParentID:       &parentID,
		Level:          int(req.Data.Attributes.Level),
		Title:          req.Data.Attributes.Title,
		Desc:           req.Data.Attributes.Desc,
		Status:         string(models.StatusInactive),
		PublicisedBy:   userId,
		LocalCost:      req.Data.Attributes.LocalCost,
		LocalCollected: &zero,
		CreatedAt:      primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	if req.Data.Attributes.Jar != nil {
		point.Jar = &models.Jar{
			Bank:      req.Data.Attributes.Jar.Bank,
			PublicURL: req.Data.Attributes.Jar.Url,
		}
	}

	res, err := server.MongoDB.Points.New().Insert(r.Context(), point)
	if err != nil {
		if err != mongo.ErrMultipleIndexDrop {
			httpkit.RenderErr(w, problems.Conflict("Main point already exists"))
			return
		}
		log.WithError(err).Error("Failed to create point")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.Point(*res))
	return
}
