package controllers

import (
	"context"
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

// TODO this is for debug to test that only a LOGGED user can do this action
func CreateAPetToAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		issuerContext := r.Context().Value(models.HttpContextStruct{}).(models.HttpContextStruct)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		issuer := issuerContext.JwtIssuer

		var foundedUser models.User

		//TODO DOT NOT LET THIS UNCHECKED
		userID, _ := primitive.ObjectIDFromHex(issuer)
		err := userCollections.FindOne(ctx, bson.M{"_id": userID}).Decode(&foundedUser)

		//Check if the user from the issuer exists
		if err != nil {
			if err == mongo.ErrNoDocuments {
				responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
				return
			}
		}
		responses.EncodeResponse(rw, http.StatusCreated, "success", map[string]interface{}{"data": userID})

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
