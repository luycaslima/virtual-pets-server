package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	What a user can do?
	* Log in / Logout
	* Access the functions of the game
		*Edit,Create and set pets in the Vivarium
		*Create a pet or catch it
		*All actions with the pet
		* ALL FUNCTIONS IN GAME
*/

type UserID primitive.ObjectID

// TODO when get the user profile not return with email nor the password
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password []byte             `json:"-" validate:"required"` //the - is for not be returned
	Pets     []PetID            `json:"pets"`
	Vivarium []VivariumID       `json:"vivarium"`
	//Inventario
	//Personagem aqui
}

type Credentials struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password []byte `json:"password,omitempty" validate:"required"`
}
