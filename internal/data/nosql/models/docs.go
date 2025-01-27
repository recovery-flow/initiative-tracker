package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Docs struct {
	ID        primitive.ObjectID `bson:"id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Desc      string             `bson:"desc" json:"desc"`
	Link      string             `bson:"link" json:"link"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
