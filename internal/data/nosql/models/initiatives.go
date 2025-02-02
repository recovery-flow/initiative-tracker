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
	Type     IniType            `bson:"type" json:"type"`
	Status   IniStatus          `bson:"status" json:"status"`

	ChatID        primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	Organizations []OrgMember        `bson:"organizations,omitempty" json:"organizations,omitempty"`

	FinalCost    int     `bson:"final_cost" json:"final_cost"`
	CollectedSum int     `bson:"collected_cost" json:"collected_cost"`
	Wallets      Wallets `bson:"wallets" json:"wallets"`

	Likes   int `bson:"likes" json:"likes"`
	Reposts int `bson:"reposts" json:"reposts"`
	Reports int `bson:"reports" json:"reports"`

	StartAt   *primitive.DateTime `bson:"start_at,omitempty" json:"start_at,omitempty"`
	EndAt     *primitive.DateTime `bson:"end_at,omitempty" json:"end_at,omitempty"`
	UpdatedAt *primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	ClosedAt  *primitive.DateTime `bson:"closed_at,omitempty" json:"closed_at,omitempty"`
}

type OrgMember struct {
	ID     primitive.ObjectID `bson:"organization_id" json:"organization_id"`
	Status StatusOrg          `bson:"status" json:"status"`
	Since  primitive.DateTime `bson:"since" json:"since"`
}

type IniStatus string

const (
	StatusDiscus    IniStatus = "discus"
	StatusPlanning  IniStatus = "planning"
	StatusActive    IniStatus = "active"
	StatusInactive  IniStatus = "inactive"
	StatusClosed    IniStatus = "closed"
	StatusPostponed IniStatus = "postponed"
	StatusDismissed IniStatus = "dismissed"
)

func GetIniStatus(s string) *IniStatus {
	status := map[string]IniStatus{
		"discus":    StatusDiscus,
		"planning":  StatusPlanning,
		"active":    StatusActive,
		"inactive":  StatusInactive,
		"closed":    StatusClosed,
		"postponed": StatusPostponed,
		"dismissed": StatusDismissed,
	}
	if iniStatus, exists := status[s]; exists {
		return &iniStatus
	}
	return nil
}

type StatusOrg string

const (
	StatusOrgFounder    StatusOrg = "founder"
	StatusOrgGeneral    StatusOrg = "general"
	StatusOrgStrategy   StatusOrg = "strategy"
	StatusOrgPrincipal  StatusOrg = "principal"
	StatusOrgAssociated StatusOrg = "associated"
)

func StatusOrgFromString(s string) *StatusOrg {
	status := map[string]StatusOrg{
		"founder":    StatusOrgFounder,
		"general":    StatusOrgGeneral,
		"strategy":   StatusOrgStrategy,
		"principal":  StatusOrgPrincipal,
		"associated": StatusOrgAssociated,
	}
	if statusOrg, exists := status[s]; exists {
		return &statusOrg
	}
	return nil
}

type IniType string

const (
	IniTypeBuilding     IniType = "building"
	IniTypeHumanitarian IniType = "humanitarian"
	IniTypePetition     IniType = "petition"
	IniTypeProtest      IniType = "protest"
	IniTypeEcology      IniType = "ecology"
	IniTypeEducation    IniType = "education"
	IniTypeHealth       IniType = "health"
	IniTypeCulture      IniType = "culture"
	IniTypeSport        IniType = "sport"
	IniTypeScience      IniType = "science"
	IniTypeCivilWork    IniType = "civil_work"
	IniTypeOther        IniType = "other"
)

func GetIniType(input string) *IniType {
	types := map[string]IniType{
		"building":     IniTypeBuilding,
		"humanitarian": IniTypeHumanitarian,
		"petition":     IniTypePetition,
		"protest":      IniTypeProtest,
		"ecology":      IniTypeEcology,
		"education":    IniTypeEducation,
		"health":       IniTypeHealth,
		"culture":      IniTypeCulture,
		"sport":        IniTypeSport,
		"science":      IniTypeScience,
		"civil_work":   IniTypeCivilWork,
		"other":        IniTypeOther,
	}

	if iniType, exists := types[input]; exists {
		return &iniType
	}
	return nil
}
