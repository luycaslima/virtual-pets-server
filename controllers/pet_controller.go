package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		//TODO change the validate JWT function to this uunder
		cookie, err := r.Cookie("jwt")
		//Check if there is a cookie
		if err != nil {
			if err == http.ErrNoCookie {
				responses.EncodeResponse(rw, http.StatusUnauthorized, "error", map[string]interface{}{"data": err.Error()})
				return
			}
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		tknStr := cookie.Value

		token, err := jwt.ParseWithClaims(tknStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				responses.EncodeResponse(rw, http.StatusUnauthorized, "error", map[string]interface{}{"data": err.Error()})
				return
			}
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
		}

		if !token.Valid {
			responses.EncodeResponse(rw, http.StatusUnauthorized, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//Find the user by the jwt token
		claims := token.Claims.(jwt.MapClaims)
		var foundedUser models.User

		//TODO DOT NOOT LET THIS UNCHECKED
		issuer, _ := claims.GetIssuer()
		userID, _ := primitive.ObjectIDFromHex(issuer)

		userCollections.FindOne(ctx, bson.M{"_id": userID}).Decode(&foundedUser)

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
