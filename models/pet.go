package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO for "wild pets" that are not owned
type WildPet struct {
}

type Pet struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" example:"1"`
	PublicData   PetPublicData      `json:"publicdata"`
	InternalData PetInternalData    `json:"internaldata"`
}

type PetPublicData struct {
	PetName    string         `json:"petname"`
	Birthday   string         `json:"birthday"`
	Techniques []string       `json:"techniques"` //TODO change to an array of struct Technique in the future
	Status     Status         `json:"status"`
	Specie     SpecieOverview `json:"specie"`
}

type PetInternalData struct {
	OwnerID   primitive.ObjectID `json:"ownerID"`
	Hunger    int                `json:"hunger" default:"0" validate:"required,min=0,max=100"`
	Happiness int                `json:"happiness" default:"100" validate:"required,min=0,max=100"`
	Cleanness int                `json:"cleanness" default:"100" validate:"required,min=0,max=100"`
}

func (p *Pet) TrainPet() bool {
	return false
}

func (p *Pet) FeedPet() {

}

func (p *Pet) PlayWithPet() {

}
