package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (user *User) AddPet(petID PetID) []PetID {
	user.Pets = append(user.Pets, petID)
	return user.Pets
}

func (user *User) AddVivarium(vivariumID VivariumID) []VivariumID {
	user.Vivarium = append(user.Vivarium, vivariumID)
	return user.Vivarium
}
