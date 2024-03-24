package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"github.com/luycaslima/virtual-pets-server/usecase"
)

func CreateASpecie(useCase *usecase.SpecieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var specie models.Specie

		if err := json.NewDecoder(r.Body).Decode(&specie); err != nil {
			responses.EncodeResponse(w, http.StatusBadRequest, "Invalid JSON body", false, err.Error())
			return
		}

		if validationErr := validate.Struct(&specie); validationErr != nil {
			errors := validationErr.(validator.ValidationErrors)
			responses.EncodeResponse(w, http.StatusBadRequest, "Validation Error", false, errors.Error())
			return
		}

		err := useCase.RegisterANewSpecie(ctx, &specie)
		if err != nil {
			responses.EncodeResponse(w, http.StatusInternalServerError, "Error inserting on the Database", false, err.Error())
			return
		}

		responses.EncodeResponse(w, http.StatusOK, "Specie Created", true, nil)
	}
}

func GetListOfSpecies(useCase *usecase.SpecieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := useCase.GetListOfSpecies(ctx)
		if err != nil {
			responses.EncodeResponse(w, http.StatusInternalServerError, "Error searching in the database", false, err.Error())
		}

		responses.EncodeResponse(w, http.StatusOK, "Success", true, results)
	}
}

func GetASpecieDetails(useCase *usecase.SpecieService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		params := mux.Vars(r)
		specieId := params["id"]

		foundedSpecie, msg, statusCode, err := useCase.GetSpecieDetails(ctx, specieId)
		if err != nil {
			responses.EncodeResponse(w, statusCode, msg, false, err.Error())
		}
		responses.EncodeResponse(w, statusCode, msg, true, foundedSpecie)
	}
}

/*

func GetListOfInitialPhaseSpecies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO implement
	}
}

func UpdateASpecie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		params := mux.Vars(r)
		specieId := params["id"]
		var foundedSpecie models.Specie

		id, err := primitive.ObjectIDFromHex(specieId)
		if err != nil {
			responses.EncodeResponse(w, http.StatusInternalServerError, "Invalid Species ID", false, err.Error())
			return
		}

		if err := speciesCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundedSpecie); err != nil {
			responses.EncodeResponse(w, http.StatusNotFound, "Specie Not Found", false, err.Error())
			return
		}

		//TODO check body and Update

	}
}
*/
