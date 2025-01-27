package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID     primitive.ObjectID `bson:"tag_id" json:"tag_id"`
	Name   string             `bson:"name" json:"name"`
	Weight int                `bson:"weight" json:"weight"`
}
