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

	res := resources.Initiative{
		Data: resources.InitiativeData{
			Id:   ini.ID.Hex(),
			Type: resources.InitiativeType,
			Attributes: resources.InitiativeDataAttributes{
				Name:      ini.Name,
				Desc:      ini.Desc,
				Goal:      ini.Goal,
				Verified:  verIni,
				Status:    string(ini.Status),
				Likes:     int64(ini.Likes),
				Reposts:   int64(ini.Reposts),
				Reports:   int64(ini.Reports),
				CreatedAt: ini.CreatedAt.Time().UTC(),
			},
			Links: resources.LinkSelf{
				Self: resources.LinkInitiativeBase + resources.LinkPrivateInitiative + ini.ID.Hex(),
			},
			Relationships: resources.InitiativeDataRelationships{
				Chat: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkChatStorageBase + resources.LinkPrivateChat + ini.ChatID.Hex(),
						Related: resources.LinkChatStorageBase + resources.LinkGetChat + ini.ChatID.Hex(),
					},
				},
				Likes: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkReactionsStorageBase + resources.LinkPrivateReactions + resources.Likes + ini.ID.Hex(),
						Related: resources.LinkReactionsStorageBase + resources.LinkGetReactions + resources.Likes + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
				Reposts: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkReactionsStorageBase + resources.LinkPrivateReactions + resources.Reposts + ini.ID.Hex(),
						Related: resources.LinkReactionsStorageBase + resources.LinkGetReactions + resources.Reposts + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
				Reports: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkReactionsStorageBase + resources.LinkPrivateReactions + resources.Reports + ini.ID.Hex(),
						Related: resources.LinkReactionsStorageBase + resources.LinkGetReactions + resources.Reports + ini.ID.Hex() + resources.Pagination10EndLink,
					},
				},
				Participants: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkInitiativeBase + resources.LinkPrivateInitiative + ini.ID.Hex() + resources.LinkParticipants,
						Related: resources.LinkInitiativeBase + resources.LinkGetInitiative + ini.ID.Hex() + resources.LinkParticipants + resources.Pagination10EndLink,
					},
				},
				Points: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Self:    resources.LinkInitiativeBase + resources.LinkPrivateInitiative + ini.ID.Hex() + resources.LinkPoints,
						Related: resources.LinkInitiativeBase + resources.LinkGetInitiative + ini.ID.Hex() + resources.LinkPoints + resources.Pagination10EndLink,
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

	return res
}
