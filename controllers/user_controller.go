package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/dto"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"github.com/luycaslima/virtual-pets-server/usecase"
)

func RegisterAnUser(useCase *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var userRegistration dto.UserRegistrationRequest

		//Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&userRegistration); err != nil {
			responses.EncodeResponse(w, http.StatusBadRequest, "Invalid JSON body", false, err.Error())
			return
		}

		//Validate the constructed struct
		if validationErr := validate.Struct(&userRegistration); validationErr != nil {
			errors := validationErr.(validator.ValidationErrors)
			responses.EncodeResponse(w, http.StatusBadRequest, "Validation Error", false, errors.Error())
			return
		}

		//TODO check if the username exist or the email is being used (in the frontend?)
		err := useCase.RegisterAnUser(ctx, &userRegistration)
		if err != nil {
			responses.EncodeResponse(w, http.StatusInternalServerError, "Error on the Database Service", false, err.Error())
			return
		}

		responses.EncodeResponse(w, http.StatusCreated, "User created!", true, nil)

	}
}

func LoginAnUser(useCase *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var userCredentials dto.UserLoginRequest

		if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
			responses.EncodeResponse(w, http.StatusBadRequest, "Invalid JSON body", false, err.Error())
			return
		}

		//Validate the constructed struct
		if validationErr := validate.Struct(&userCredentials); validationErr != nil {
			errors := validationErr.(validator.ValidationErrors)
			responses.EncodeResponse(w, http.StatusBadRequest, "Validation Error", false, errors.Error())
			return
		}

		cookie, msg, statusCode, err := useCase.LoginAnUser(ctx, &userCredentials)
		//TODO pass the message too?
		if err != nil {
			responses.EncodeResponse(w, statusCode, msg, false, err.Error())
			return
		}

		http.SetCookie(w, cookie)

		responses.EncodeResponse(w, statusCode, msg, true, nil)
	}
}

func GetAnUserProfile(useCase *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var params = mux.Vars(r)
		username := params["username"]

		userPublicData, err := useCase.GetUserPublicDataByUsername(ctx, username)
		if err != nil {
			responses.EncodeResponse(w, http.StatusNotFound, "User not found!", false, err.Error())
			return
		}
		responses.EncodeResponse(w, http.StatusOK, "User Founded!", true, userPublicData)
	}
}

func AdoptAPet(useCase *usecase.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		issuerContext := r.Context().Value(models.HttpContextStruct{}).(models.HttpContextStruct)
		issuer := issuerContext.JwtIssuer

		var petRequest dto.CreatePetRequest

		//Validate body
		if err := json.NewDecoder(r.Body).Decode(&petRequest); err != nil {
			responses.EncodeResponse(w, http.StatusBadRequest, "Invalid Json Body", false, err.Error())
			return
		}

		if validationErr := validate.Struct(&petRequest); validationErr != nil {
			responses.EncodeResponse(w, http.StatusBadRequest, "Validation error", false, validationErr.Error())
			return
		}

		newPet, msg, statusCode, err := useCase.CreateAPet(ctx, issuer, &petRequest)
		if err != nil {
			responses.EncodeResponse(w, statusCode, msg, false, err.Error())
			return
		}

		responses.EncodeResponse(w, statusCode, msg, true, newPet)
	}
}
