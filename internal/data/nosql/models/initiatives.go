package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Initiatives struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	OwnerID   primitive.ObjectID  `bson:"owner_id" json:"owner_id"`
	ProjectID primitive.ObjectID  `bson:"project_id" json:"project_id"`
	WalletID  *primitive.ObjectID `bson:"wallet_id,omitempty" json:"wallet_id,omitempty"`

	Name string `bson:"name" json:"name"`
	Desc string `bson:"desc" json:"desc"`
	Goal string `bson:"goal" json:"goal"`

	Status InitiativeStatus `bson:"status" json:"status"`

	FinalCost    *int `bson:"final_cost,omitempty" json:"final_cost,omitempty"`
	CollectedSum *int `bson:"collected_cost,omitempty" json:"collected_cost,omitempty"`

	Likes   int `bson:"likes" json:"likes"`
	Reports int `bson:"reports" json:"reports"`

	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	ClosedAt  *primitive.DateTime `bson:"closed_at,omitempty" json:"closed_at,omitempty"`
	CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`
}

type InitiativeStatus string

const (
	InitiativeStatusActive    InitiativeStatus = "active"
	InitiativeStatusInactive  InitiativeStatus = "inactive"
	InitiativeStatusClosed    InitiativeStatus = "closed"
	InitiativeStatusDismissed InitiativeStatus = "dismissed"
)

func GetInitiativeStatus(s string) *InitiativeStatus {
	status := map[string]InitiativeStatus{
		"active":    InitiativeStatusActive,
		"inactive":  InitiativeStatusInactive,
		"closed":    InitiativeStatusClosed,
		"dismissed": InitiativeStatusDismissed,
	}
	if iniStatus, exists := status[s]; exists {
		return &iniStatus
	}
	return nil
}
