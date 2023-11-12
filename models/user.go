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

// User represents the model  for an User's Data, password is ALWAYS HASHED
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty" example:"1"`
	Username string             `json:"username,omitempty" validate:"required" example:"ronald123"`
	Email    string             `json:"email,omitempty" validate:"required" example:"ronald@email.com"`
	Password string             `json:"password" validate:"required" example:"aksdmalknj@154/JKNJ"`
	Pets     []PetID            `json:"pets" `
	Vivarium []VivariumID       `json:"vivarium"`
	Money    int32              `json:"money" example:"500000"`
	//Inventario
	//Personagem aqui
}

// Credentials represent the body of Request of an User Loggin
type UserCredentials struct {
	Username string `json:"username,omitempty" validate:"required" example:"ronald123"`
	Password string `json:"password,omitempty" validate:"required" example:"aksdmalknj@154/JKNJ"`
}
