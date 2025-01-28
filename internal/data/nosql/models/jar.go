package models

type Jar struct {
	Bank      string `bson:"bank" json:"bank"`
	PublicURL string `bson:"public_url" json:"public_url"`
}
