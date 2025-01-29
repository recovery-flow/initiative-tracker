package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Initiative struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Desc     string             `bson:"desc" json:"desc"`
	Goal     string             `bson:"goal" json:"goal"`
	Verified bool               `bson:"verified" json:"verified"`
	Location *string            `bson:"location,omitempty" json:"location,omitempty"`
	Status   Status             `bson:"status" json:"status"`
	Tags     []Tag              `bson:"tags,omitempty" json:"tags,omitempty"`
	ChatID   primitive.ObjectID `bson:"chat_id" json:"chat_id"`

	Likes   int `bson:"likes" json:"likes"`
	Reposts int `bson:"reposts" json:"reposts"`
	Reports int `bson:"reports" json:"reports"`

	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	ClosedAt  *primitive.DateTime `bson:"closed_at,omitempty" json:"closed_at,omitempty"`
}

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusClosed   Status = "closed"
)

func StatusFromString(s string) *Status {
	switch s {
	case "active":
		status := StatusActive
		return &status
	case "inactive":
		status := StatusInactive
		return &status
	case "closed":
		status := StatusClosed
		return &status
	default:
		return nil
	}
}

//TODO add tags to the model this model is example

type Tag struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Weight int                `bson:"weight" json:"weight"`
}
