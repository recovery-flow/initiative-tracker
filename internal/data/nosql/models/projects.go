package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Projects struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	OwnerID  primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	ServerID primitive.ObjectID `bson:"server_id" json:"server_id"`
	WalletID primitive.ObjectID `bson:"wallet_id" json:"wallet_id"`

	Name     string `bson:"name" json:"name"`
	Desc     string `bson:"desc" json:"desc"`
	Goal     string `bson:"goal" json:"goal"`
	Verified bool   `bson:"verified" json:"verified"`

	Status ProjectStatus `bson:"status" json:"status"`

	FinalCost    *int `bson:"final_cost,omitempty" json:"final_cost,omitempty"`
	CollectedSum *int `bson:"collected_cost,omitempty" json:"collected_cost,omitempty"`

	Likes   int `bson:"likes" json:"likes"`
	Reports int `bson:"reports" json:"reports"`

	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	ClosedAt  *primitive.DateTime `bson:"closed_at,omitempty" json:"closed_at,omitempty"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
}

type ProjectStatus string

const (
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusInactive  ProjectStatus = "inactive"
	ProjectStatusClosed    ProjectStatus = "closed"
	ProjectStatusDismissed ProjectStatus = "dismissed"
)

func GetProjectStatus(s string) *ProjectStatus {
	status := map[string]ProjectStatus{
		"active":    ProjectStatusActive,
		"inactive":  ProjectStatusInactive,
		"closed":    ProjectStatusClosed,
		"dismissed": ProjectStatusDismissed,
	}
	if iniStatus, exists := status[s]; exists {
		return &iniStatus
	}
	return nil
}
