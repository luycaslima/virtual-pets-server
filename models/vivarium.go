package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VivariumID primitive.ObjectID

type Vivarium struct {
	ID         primitive.ObjectID
	PetsInside []PetID
}
