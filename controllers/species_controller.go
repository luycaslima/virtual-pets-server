package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var speciesCollection *mongo.Collection = configs.GetCollection(configs.DB, "species")

func CreateASpecie() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		var specie models.Specie
		defer cancel()

		if err := json.NewDecoder(r.Body).Decode(&specie); err != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		if validationErr := validate.Struct(&specie); validationErr != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": validationErr.Error()})
			return
		}

		newSpecie := models.Specie{
			ID:                  primitive.NewObjectID(),
			Name:                specie.Name,
			BaseStatus:          specie.BaseStatus,
			LearnableTechniques: specie.LearnableTechniques,
		}

		result, err := speciesCollection.InsertOne(ctx, newSpecie)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusCreated, "success", map[string]interface{}{"data": result})
	}
}
