package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"go.mongodb.org/mongo-driver/bson"
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

// TODO create a type form on the species
func GetBabyFormSpecies() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//TODO set cache control in all GETTERS
		rw.Header().Set("Cache-Control", "max-age=3600")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		var results []models.Specie

		//find records
		//pass these options to the Find method
		//findOptions := options.Find()
		//Set the limit of the number of record to find
		//findOptions.SetLimit(5)

		cursor, err := speciesCollection.Find(ctx, bson.M{} /*,findOptions*/)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//TODO study context and TODO
		//could use cursor.All and decode all founded directly
		for cursor.Next(ctx) {
			var elem models.Specie
			err := cursor.Decode(&elem)
			if err != nil {
				responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
				//log.Fatal(err.Error())
				return
			}
			results = append(results, elem)
		}

		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		//TODO whhy close?
		//Close the cursor once finished
		cursor.Close(context.TODO())

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": results})
	}

}
