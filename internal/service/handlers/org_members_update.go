package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/internal/service/requests"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitiativeUpdateOrgsMembers(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	req, err := requests.NewOrgMembersUpdate(r)
	if err != nil {
		log.Info("Failed to parse request: ", err)
		httpkit.RenderErr(w, problems.BadRequest(err)...)
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

	//TODO: add verification for this one as well
	initiatorId, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	status := models.GetStatusOrg(req.Data.Attributes.Status)
	if status == nil {
		log.Info("Failed to parse status")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/attributes/status": validation.NewError("status", "Invalid status"),
		})...)
		return
	}

	orgId, err := primitive.ObjectIDFromHex(req.Data.Attributes.OrgId)
	if err != nil {
		log.WithError(err).Error("Failed to parse organization id")
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/attributes/org_id": validation.NewError("org_id", "Invalid organization id"),
		})...)
		return
	}

	_, err = server.MongoDB.Initiative.New().FilterExact(map[string]any{
		"_id": iniId,
	}).AddOrgMember(r.Context(), models.OrgMember{
		ID:     orgId,
		Status: *status,
		Since:  primitive.NewDateTimeFromTime(time.Now().UTC()),
	})

	log.Infof("Wallets for initiative %s updated by %s", iniId, initiatorId)
	httpkit.Render(w, http.StatusAccepted)
}
