package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserID primitive.ObjectID

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Pets     []PetID            `json:"pets"`
	Vivarium []VivariumID       `json:"vivarium"`
	//Inventario
	//Personagem aqui
}
