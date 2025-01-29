package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Point struct {
	ID           primitive.ObjectID  `bson:"_id" json:"id"`
	InitiativeID primitive.ObjectID  `bson:"initiative_id" json:"initiative_id"`
	ParentID     *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	Level        int                 `bson:"level" json:"level"`

	Title        string    `bson:"title" json:"title"`
	Desc         string    `bson:"desc" json:"desc"`
	PublicisedBy uuid.UUID `bson:"publicised_by" json:"publicised_by"`

	LocalCost      *float64 `bson:"local_cost,omitempty" json:"local_cost,omitempty"`
	LocalCollected *float64 `bson:"local_collected,omitempty" json:"local_collected,omitempty"`
	Jar            *Jar     `bson:"bank_info,omitempty" json:"bank_info,omitempty"`

	Status string `bson:"status" json:"status"`
	Docs   []Docs `bson:"docs,omitempty" json:"docs,omitempty"`

	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
