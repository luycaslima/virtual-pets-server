package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/auth"
	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/models"
	"github.com/luycaslima/virtual-pets-server/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollections *mongo.Collection = configs.GetCollection(configs.DB, "users")

// RegisterAUser godoc
//
//	@Summary		Register a new User
//	@Description	Register a new User with a Username, Email and Password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true	"Register a new User"
//	@Success		201		{object}	models.User
//	@Router			/users/register [post]
func RegisterAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		var user models.User //Recieved user from JSON
		defer cancel()

		//validate request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//Use the validator to validate the required struct
		if validationErr := validate.Struct(&user); validationErr != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": validationErr.Error()})
			return
		}

		//TODO Understand better the hashFunction. why cost 14?
		//Encrypt password
		password, _ := auth.HashPassword(string(user.Password))

		//Data to be sended to the server
		newUser := models.User{
			ID:       primitive.NewObjectID(),
			Username: user.Username,
			Email:    user.Email,
			Password: password,
			Pets:     make([]models.PetID, 0),
			Vivarium: make([]models.VivariumID, 0),
			Money:    100,
		}

		result, err := userCollections.InsertOne(ctx, newUser)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusCreated, "success", map[string]interface{}{"data": result})
	}
}

func LoginAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		var user models.UserCredentials
		defer cancel()

		//validate request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//Use the validator to validate the required struct
		if validationErr := validate.Struct(&user); validationErr != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": validationErr.Error()})
			return
		}

		var foundedUser models.User
		//Find Username
		err := userCollections.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundedUser)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//Check password
		if isPasswordCorrect := auth.CheckPassword(user.Password, foundedUser.Password); !isPasswordCorrect {
			responses.EncodeResponse(rw, http.StatusUnauthorized, "error", map[string]interface{}{"data": "wrong password"})
			return
		}

		//return a created JWT TOKEN
		token, expirationDate, err := auth.CreateJWT(foundedUser.ID.Hex())

		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		//Set cookie of the logged session
		cookie := http.Cookie{
			SameSite: http.SameSiteNoneMode, //when the api and the site is on diferent DOMAINS
			Name:     "jwt",
			Value:    token,
			Expires:  expirationDate,
			HttpOnly: true, //to the frontend no have access
		}

		http.SetCookie(rw, &cookie)

		//return cookie with jwt
		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": "Logged with success"})
	}
}

//TODO have a renew token method

func LogoutAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		//Set cookie minus one hour to invalidate the session
		cookie := http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(rw, &cookie)

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": "Logout successfuly!"})
	}
}

func GetAUsersProfile() http.HandlerFunc {
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
			responses.EncodeResponse(rw, http.StatusBadRequest, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"user": user})
	}
}

func CheckAuthenticatedUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		defer cancel()
		cookie, err := r.Cookie("jwt")

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

		responses.EncodeResponse(rw, http.StatusFound, "success", map[string]interface{}{"data": foundedUser})

	}
}

func ChangePasswordOfAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
