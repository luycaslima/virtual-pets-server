package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var petCollection *mongo.Collection = configs.GetCollection(configs.DB, "pets")

func CreateAPet() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		issuerContext := r.Context().Value(models.HttpContextStruct{}).(models.HttpContextStruct)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		var pet models.Pet
		defer cancel()

		issuer := issuerContext.JwtIssuer
		var foundedUser models.User
		var foundedSpecie models.Specie
		var newPet models.Pet
		//Validate body
		err := json.NewDecoder(r.Body).Decode(&pet)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}
		if validationErr := validate.Struct(&pet); validationErr != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": validationErr.Error()})
		}

		//Check if the user exists
		userID, _ := primitive.ObjectIDFromHex(issuer)
		err = userCollections.FindOne(ctx, bson.M{"_id": userID}).Decode(&foundedUser)

		//Check if the user from the issuer exists
		if err != nil {
			if err == mongo.ErrNoDocuments {
				responses.EncodeResponse(rw, http.StatusNotFound, "error", map[string]interface{}{"data": err.Error()})
				return
			}
		}

		//Check if the specie exists
		err = speciesCollection.FindOne(ctx, bson.M{"_id": pet.SpecieID}).Decode(&foundedSpecie)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
				return
			}
		}

		newPet = models.Pet{
			ID:        primitive.NewObjectID(),
			PetName:   pet.PetName,
			SpecieID:  pet.SpecieID,
			OwnerID:   userID,
			Status:    foundedSpecie.BaseStatus,
			Birthday:  time.Now().String(),
			Happiness: 100,
			Hunger:    0,
			Cleanness: 100,
			//Add techniques
		}

		_, err = petCollection.InsertOne(ctx, newPet)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusCreated, "success", map[string]interface{}{"data": newPet.ID})
	}
}

func GenerateAPet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Pass here which train
func TrainPet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
