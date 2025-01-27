package models

type Jar struct {
	BankName  string `bson:"bank_name" json:"bank_name"`
	PublicURL string `bson:"public_url" json:"public_url"`
}
