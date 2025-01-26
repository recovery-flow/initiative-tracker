package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tags struct {
	ID          primitive.ObjectID   `bson:"_id" json:"id"`
	Name        string               `bson:"name" json:"name"`
	Description string               `bson:"description" json:"description"`
	UsageCount  int                  `bson:"usage_count" json:"usage_count"`
	Weight      float32              `bson:"weight" json:"weight"`
	RelatedTags []primitive.ObjectID `bson:"related_tags" json:"related_tags"`
	CreatedAt   time.Time            `bson:"created_at" json:"created_at"`
}
