package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PetID primitive.ObjectID

// TODO for pets that are not owned
type WildPet struct {
}

type Pet struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	PetName   string             `json:"petName" validate:"required"`
	Hunger    int                `json:"hunger" default:"0"`
	Happiness int                `json:"happiness" default:"100"`
	Cleanness int                `json:"cleanness" default:"100"`
	Status    Status             `json:"status"`
	OwnerID   UserID             `json:"ownerID" validate:"required"`
	Birthday  primitive.DateTime `json:"birthday"`
}
