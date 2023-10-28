package controllers

import (
	"net/http"

	"github.com/luycaslima/virtual-pets-server/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

var petCollection *mongo.Collection = configs.GetCollection(configs.DB, "pets")

// Pass here which train
func TrainPet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
