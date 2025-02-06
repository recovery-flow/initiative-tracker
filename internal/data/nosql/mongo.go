package nosql

import (
	"github.com/recovery-flow/initiative-tracker/internal/data/nosql/repositories"
)

type Repo struct {
	Initiative repositories.Initiatives
}

func NewRepositoryNoSql(uri, dbName string) (*Repo, error) {
	initiativeRepo, err := repositories.NewInitiative(uri, dbName, "initiatives")
	if err != nil {
		return nil, err
	}
	return &Repo{
		Initiative: initiativeRepo,
	}, nil
}
