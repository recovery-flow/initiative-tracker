package nosql

import (
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/models"
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/repositories"
)

type Repo struct {
	Initiative   repositories.Iniative
	Participants repositories.Participant
	Tags         repositories.Tag
}

func NewRepositoryNoSql(uri, dbName string) (*Repo, error) {
	initiativeRepo, err := repositories.NewInitiative(uri, dbName, models.InitiativeCollection)
	if err != nil {
		return nil, err
	}
	participantRepo, err := repositories.NewParticipant(uri, dbName, models.ParticipantCollection)
	if err != nil {
		return nil, err
	}
	tagRepo, err := repositories.NewTag(uri, dbName, models.TagCollection)
	if err != nil {
		return nil, err
	}
	return &Repo{
		Initiative:   initiativeRepo,
		Participants: participantRepo,
		Tags:         tagRepo,
	}, nil
}
