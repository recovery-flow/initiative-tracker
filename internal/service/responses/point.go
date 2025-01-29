package responses

import (
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/resources"
)

func Point(point models.Point) resources.Point {
	var res resources.Point
	if point.ParentID != nil {
		parent := point.ParentID.Hex()
		res.Data.Attributes.ParentId = &parent

		res.Data.Relationships.Parent = &resources.LinksDirect{
			Links: resources.LinksDirectLinks{
				Related: resources.BaseInitiative + resources.InitiativeEndpoints.Base.Public + point.InitiativeID.Hex() + resources.InitiativeEndpoints.Points.Public + point.ParentID.Hex(),
				Self:    resources.BaseInitiative + resources.InitiativeEndpoints.Base.Private + point.InitiativeID.Hex() + resources.InitiativeEndpoints.Points.Private + point.ParentID.Hex(),
			},
			Data: &resources.LinksDirectData{
				Id:   point.ParentID.Hex(),
				Type: resources.PointType,
			},
		}
	}

	res = resources.Point{
		Data: resources.PointData{
			Id:   point.ID.Hex(),
			Type: resources.PointType,
			Attributes: resources.PointAttributes{
				InitiativeId:   point.InitiativeID.Hex(),
				Level:          int32(point.Level),
				Title:          point.Title,
				Desc:           point.Desc,
				PublishedBy:    point.PublicisedBy.String(),
				LocalCost:      *point.LocalCost,
				LocalCollected: *point.LocalCollected,
				Status:         point.Status,
				CreatedAt:      point.CreatedAt.Time().UTC(),
			},
			Relationships: resources.PointDataRelationships{
				Initiative: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Related: resources.BaseInitiative + resources.InitiativeEndpoints.Base.Public + point.InitiativeID.Hex(),
						Self:    resources.BaseInitiative + resources.InitiativeEndpoints.Base.Private + point.InitiativeID.Hex(),
					},
					Data: &resources.LinksDirectData{
						Id:   point.InitiativeID.Hex(),
						Type: resources.InitiativeType,
					},
				},
				Plan: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Related: resources.BaseInitiative + resources.InitiativeEndpoints.Base.Public + point.InitiativeID.Hex() + resources.InitiativeEndpoints.Points.Public,
						Self:    resources.BaseInitiative + resources.InitiativeEndpoints.Base.Private + point.InitiativeID.Hex() + resources.InitiativeEndpoints.Points.Private,
					},
					Data: &resources.LinksDirectData{
						Id:   point.InitiativeID.Hex(),
						Type: resources.PlanType,
					},
				},
				PublishedBy: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Related: resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Public + point.PublicisedBy.String(),
						Self:    resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Private + point.PublicisedBy.String(),
					},
					Data: &resources.LinksDirectData{
						Id:   point.PublicisedBy.String(),
						Type: "user",
					},
				},
			},
		},
	}

	if point.UpdatedAt != nil {
		timeNow := point.UpdatedAt.Time().UTC()
		res.Data.Attributes.UpdatedAt = &timeNow
	}

	if point.Jar != nil {

		jar := resources.JarData{
			Id:   point.ID.Hex(),
			Type: resources.JarType,
			Attributes: resources.JarAttributes{
				Bank: point.Jar.Bank,
				Url:  point.Jar.PublicURL,
			},
		}

		res.Included = append(res.Included, jar)
	}

	return res
}
