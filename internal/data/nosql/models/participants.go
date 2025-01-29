package models

import (
	"github.com/google/uuid"
	"github.com/recovery-flow/roles"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	UserID       uuid.UUID          `bson:"user_id" json:"user_id"`
	InitiativeID primitive.ObjectID `bson:"initiative_id" json:"initiative_id"`
	FirstName    string             `bson:"first_name" json:"first_name"`
	SecondName   string             `bson:"second_name" json:"second_name"`
	ThirdName    *string            `bson:"third_name,omitempty" json:"third_name,omitempty"`
	DisplayName  string             `bson:"display_name" json:"display_name"`
	Position     string             `bson:"position" json:"position"`
	Desc         *string            `bson:"desc,omitempty" json:"desc,omitempty"`
	Verified     bool               `bson:"verified" json:"verified"`
	Role         roles.TeamRole     `bson:"role" json:"role"`

	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
