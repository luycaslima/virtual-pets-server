package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PetID primitive.ObjectID

type Pet struct {
	ID        primitive.ObjectID
	PetName   string
	Hunger    int
	Happiness int
	Cleaness  int
	Status    Status
	BornAt    primitive.DateTime
}
