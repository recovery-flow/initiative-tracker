package responses

import (
	"net/url"

	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/resources"
)

func PointsCollection(points []models.Point, baseURL string, queryParams url.Values, pageSize, pageNumber int64) resources.PointsCollection {
	var res []resources.PointData
	for _, el := range points {
		res = append(res, Point(el).Data)
	}

	links := resources.LinksPagination{
		Self:     *generatePaginationLink(baseURL, queryParams, pageNumber, pageSize),
		Previous: *generatePaginationLink(baseURL, queryParams, pageNumber-1, pageSize),
		Next:     *generatePaginationLink(baseURL, queryParams, pageNumber+1, pageSize),
	}

	return resources.PointsCollection{
		Data:  res,
		Links: links,
	}
}
