package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//TODO create a Admin user structure

type UserPublicData struct {
	Username    string               `json:"username" example:"ronald123"`
	Pets        []primitive.ObjectID `json:"pets"`
	Vivariums   []primitive.ObjectID `json:"vivariums"`
	CreatedDate string               `json:"createddate"`
}

type UserPrivateData struct {
	Money uint   `json:"money" example:"29387"`
	Email string `json:"email" example:"test@mail.com"`
}

type UserInternalData struct {
	Password string `json:"password" example:"aksçdaksnd231@"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id" example:"18923k1j2h"`
	PublicData   UserPublicData     `json:"publicdata"`
	PrivateData  UserPrivateData    `json:"privatedata"`
	InternalData UserInternalData   `json:"internaldata"`
}

func (user *User) AddPet(petID primitive.ObjectID) []primitive.ObjectID {
	user.PublicData.Pets = append(user.PublicData.Pets, petID)
	return user.PublicData.Pets
}

func (user *User) AddVivarium(vivariumID primitive.ObjectID) []primitive.ObjectID {
	user.PublicData.Vivariums = append(user.PublicData.Vivariums, vivariumID)
	return user.PublicData.Vivariums
}
