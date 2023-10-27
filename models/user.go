package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID
	Username string
	Pets     []PetID
	Vivarium []VivariumID
}
