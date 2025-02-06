package responses

import (
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/resources"
)

func Initiative(ini models.Initiative) resources.Initiative {
	verIni := false
	if ini.Verified {
		verIni = true
	}

	finalCost := int64(ini.FinalCost)
	collectedSum := int64(ini.CollectedSum)

	res := resources.Initiative{
		Data: resources.InitiativeData{
			Id:   ini.ID.Hex(),
			Type: resources.InitiativeType,
			Attributes: resources.InitiativeDataAttributes{
				Name:         ini.Name,
				Desc:         ini.Desc,
				Goal:         ini.Goal,
				Verified:     verIni,
				Type:         string(ini.Type),
				Status:       string(ini.Status),
				FinalCost:    &finalCost,
				CollectedSum: &collectedSum,
				Likes:        int64(ini.Likes),
				Reposts:      int64(ini.Reposts),
				Reports:      int64(ini.Reports),
			},
			Links: resources.LinkSelf{
				Self: resources.BaseInitiative + resources.InitiativeEndpoints.Base.Private + ini.ID.Hex(),
			},
			Relationships: resources.InitiativeDataRelationships{
				Chat: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.BaseChatStorage + resources.ChatEndpoints.Private + ini.ChatID.Hex(),
						Related: resources.BaseChatStorage + resources.ChatEndpoints.Public + ini.ChatID.Hex(),
					},
					Data: &resources.LinksDirectData{
						Id:   ini.ChatID.Hex(),
						Type: "chat",
					},
				},
				Likes: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.BaseReactionsStorage + resources.ReactionsEndpoints.Likes.Private + ini.ID.Hex(),
						Related: resources.BaseReactionsStorage + resources.ReactionsEndpoints.Likes.Public + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
				Reposts: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.BaseReactionsStorage + resources.ReactionsEndpoints.Reposts.Private + ini.ID.Hex(),
						Related: resources.BaseReactionsStorage + resources.ReactionsEndpoints.Reposts.Public + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
				Reports: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.BaseReactionsStorage + resources.ReactionsEndpoints.Reports.Private + ini.ID.Hex(),
						Related: resources.BaseReactionsStorage + resources.ReactionsEndpoints.Reports.Public + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
			},
		},
	}
	if ini.Location != nil {
		res.Data.Attributes.Location = ini.Location
	}
	if ini.UpdatedAt != nil {
		upAt := ini.UpdatedAt.Time().UTC()
		res.Data.Attributes.UpdatedAt = &upAt
	}
	if ini.ClosedAt != nil {
		clAt := ini.ClosedAt.Time().UTC()
		res.Data.Attributes.ClosedAt = &clAt
	}
	if ini.StartAt != nil {
		stAt := ini.StartAt.Time().UTC()
		res.Data.Attributes.StartAt = &stAt
	}
	if ini.EndAt != nil {
		enAt := ini.EndAt.Time().UTC()
		res.Data.Attributes.EndAt = &enAt
	}
	return res
}
