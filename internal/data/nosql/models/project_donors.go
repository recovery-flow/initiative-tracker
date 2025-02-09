package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProjectDonor struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	ProjectID primitive.ObjectID `bson:"project_id" json:"project_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`

	Amount      int                       `bson:"amount" json:"amount"`
	CreatedAt   primitive.DateTime        `bson:"created_at" json:"created_at"`
	Initiatives []ProjectDonorInitiatives `bson:"initiatives,omitempty" json:"initiatives,omitempty"`
}

type ProjectDonorInitiatives struct {
	InitiativeID primitive.ObjectID `bson:"initiative_id" json:"initiative_id"`
	Amount       int                `bson:"amount" json:"amount"`
}
