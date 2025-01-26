package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Initiative struct {
	ID       primitive.ObjectID   `bson:"_id" json:"id"`
	Name     string               `bson:"name" json:"name"`
	Desc     string               `bson:"desc" json:"desc"`
	Goal     string               `bson:"goal" json:"goal"`
	Tags     []primitive.ObjectID `bson:"tags" json:"tags"`
	Location *string              `bson:"location,omitempty" json:"location,omitempty"`
	Team     []Participants       `bson:"team" json:"team"`
	ChatID   primitive.ObjectID   `bson:"chat_id" json:"chat_id"`
}
