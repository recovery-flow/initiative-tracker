package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParticipantsByOrganization(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}
	log := server.Logger

	queryParams := r.URL.Query()
	filterResponse := make(map[string]any)
	for key, values := range queryParams {
		if len(values) > 0 && key != "page[size]" && key != "page[number]" {
			filterResponse[key] = values[0]
		}
	}

	pageSize := 10
	pageNumber := 1

	if size := queryParams.Get("page[size]"); size != "" {
		if parsedSize, err := strconv.Atoi(size); err == nil && parsedSize > 0 {
			pageSize = parsedSize
		}
	}

	if number := queryParams.Get("page[number]"); number != "" {
		if parsedNumber, err := strconv.Atoi(number); err == nil && parsedNumber > 0 {
			pageNumber = parsedNumber
		}
	}

	limit := int64(pageSize)
	skip := int64((pageNumber - 1) * pageSize)

	iniID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "initiative_id"))
	if err != nil {
		log.WithError(err).Error("Failed to parse initiative id")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	filters := make(map[string]any)
	filters["initiative_id"] = iniID

	participants, err := server.MongoDB.Participants.New().Filter(filters).Limit(limit).Skip(skip).Select(r.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			httpkit.RenderErr(w, problems.NotFound("Participants not found"))
			return
		}
		log.WithError(err).Error("Failed to get participants")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	baseURL := "private/organization/" + iniID.Hex() + "/participant"

	log.Infof("Participants: %v", participants)
	httpkit.Render(w, responses.ParticipantCollection(participants, baseURL, queryParams, int64(pageSize), int64(pageNumber)))
}
