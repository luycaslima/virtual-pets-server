package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetID primitive.ObjectID

// TODO for pets that are not owned
type WildPet struct {
}

type Pet struct {
	ID         primitive.ObjectID `json:"_id,omitempty"`
	PetName    string             `json:"petName" validate:"required"`
	Hunger     int                `json:"hunger" default:"0"`
	Happiness  int                `json:"happiness" default:"100"`
	Cleanness  int                `json:"cleanness" default:"100"`
	SpecieID   SpecieID           `json:"specieID" validate:"required"`
	Status     Status             `json:"status"`
	OwnerID    UserID             `json:"ownerID"`
	Birthday   string             `json:"birthday"`
	Techniques []string           `json:"techniques"` //TODO change to an array of struct Technique in the future
}
