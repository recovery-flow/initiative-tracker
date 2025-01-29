package responses

import (
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/resources"
)

func Participant(participant models.Participant) resources.Participant {
	ver := false
	if participant.Verified {
		ver = true
	}
	return resources.Participant{
		Data: resources.ParticipantData{
			Id:   participant.UserID.String(),
			Type: resources.ParticipantType,
			Attributes: resources.ParticipantDataAttributes{
				FirstName:   participant.FirstName,
				SecondName:  participant.SecondName,
				ThirdName:   participant.ThirdName,
				DisplayName: participant.DisplayName,
				Position:    participant.Position,
				Verified:    ver,
				Desc:        participant.Desc,
				Role:        string(participant.Role),
				CreatedAt:   participant.CreatedAt.Time().UTC(),
			},
			Relationships: resources.ParticipantDataRelationships{
				User: resources.LinksDirect{
					Links: resources.LinksDirectLinks{
						Related: resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Public + participant.UserID.String(),
						Self:    resources.BaseUserStorage + resources.UserStorageEndpoints.Base.Private + participant.UserID.String(),
					},
				},
			},
		},
	}
}
