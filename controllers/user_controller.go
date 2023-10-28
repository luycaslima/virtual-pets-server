package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollections *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		var user models.User //Recieved user from JSON
		defer cancel()

		//validate request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//Use the validator to validate the required struct
		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//Data to be sended to the server
		newUser := models.User{
			ID:       primitive.NewObjectID(),
			Username: user.Username,
			Email:    user.Email,
			Pets:     make([]models.PetID, 0),
			Vivarium: make([]models.VivariumID, 0),
		}

		result, err := userCollections.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		var params = mux.Vars(r)
		//userID := params["userID"]
		username := params["username"]
		var user models.User
		defer cancel()

		//Convert ID
		//objID, _ := primitive.ObjectIDFromHex(userID)
		//err := userCollections.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
		err := userCollections.FindOne(ctx, bson.M{"username": username}).Decode(&user)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}
		rw.WriteHeader(http.StatusOK)
		response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": user}}
		json.NewEncoder(rw).Encode(response)
	}
}

func EditAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func CreateAPetToAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
