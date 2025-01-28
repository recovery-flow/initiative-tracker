package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Point struct {
	ID           primitive.ObjectID  `bson:"_id" json:"id"`
	InitiativeID primitive.ObjectID  `bson:"initiative_id" json:"initiative_id"`
	ParentID     *primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`

	Title               string    `bson:"title" json:"title"`
	Desc                string    `bson:"desc" json:"desc"`
	PublicisedBy        uuid.UUID `bson:"publicised_by" json:"publicised_by"`
	LocalCost           float64   `bson:"local_cost" json:"local_cost"`
	LocalCollected      float64   `bson:"local_collected" json:"local_collected"`
	CumulativeCost      float64   `bson:"cumulative_cost" json:"cumulative_cost"`
	CumulativeCollected float64   `bson:"cumulative_collected" json:"cumulative_collected"`
	Status              string    `bson:"status" json:"status"`
	Jar                 Jar       `bson:"bank_info" json:"bank_info"`
	Docs                []Docs    `bson:"docs,omitempty" json:"docs,omitempty"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
