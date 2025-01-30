package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/initiative-tracker/internal/config"
	"github.com/recovery-flow/initiative-tracker/internal/service/responses"
	"github.com/recovery-flow/initiative-tracker/resources"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	filter := bson.M{}

	// Парсинг query параметров
	queryParams := r.URL.Query()

	validFilters := map[string]string{
		"initiative_id": "initiative_id",
		"parent_id":     "parent_id",
		"publicised_by": "publicised_by",
		"id":            "_id",
	}

	for param, field := range validFilters {
		value := queryParams.Get(param)
		if value == "" {
			continue
		}

		switch param {
		case "id":
			objID, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"id": validation.NewError("id", "Invalid ID format"),
				})...)
				return
			}
			filter[field] = objID

		case "initiative_id":
			objID, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"initiative_id": validation.NewError("initiative_id", "Invalid ID format"),
				})...)
				return
			}
			filter[field] = objID

		case "parent_id":
			objID, err := primitive.ObjectIDFromHex(value)
			if err != nil {
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"parent_id": validation.NewError("parent_id", "Invalid ID format"),
				})...)
				return
			}
			filter[field] = objID

		case "publicised_by":
			_, err := uuid.Parse(value)
			if err != nil {
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"publicised_by": validation.NewError("publicised_by", "Invalid UUID format"),
				})...)
				return
			}
			filter[field] = value

		case "created_at":
			t, err := time.Parse(time.RFC3339, value)
			if err != nil {
				httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
					"created_at": validation.NewError("created_at", "Invalid date format"),
				})...)
				return
			}
			filter[field] = primitive.NewDateTimeFromTime(t)
		}
	}

	log.Infof("Filter: %v", filter)

	pageSize, pageNumber := parsePagination(queryParams)
	limit := int64(pageSize)
	skip := int64((pageNumber - 1) * pageSize)

	res, err := server.MongoDB.Points.New().
		Filter(filter).
		Limit(limit).
		Skip(skip).
		Select(r.Context())

	if err != nil {
		if err == mongo.ErrNoDocuments {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.WithError(err).Error("Failed to fetch points")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if len(res) == 0 {
		httpkit.RenderErr(w, problems.NotFound())
		return
	}

	httpkit.Render(w, responses.PointsCollection(res, resources.ServiceHost+r.URL.Path, queryParams, int64(pageSize), int64(pageNumber)))
}

func parsePagination(q url.Values) (size, number int) {
	size = 10
	number = 1

	if s := q.Get("page[size]"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			size = n
		}
	}

	if n := q.Get("page[number]"); n != "" {
		if num, err := strconv.Atoi(n); err == nil && num > 0 {
			number = num
		}
	}

	return size, number
}
