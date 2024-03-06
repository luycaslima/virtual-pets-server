package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

// RegisterAnUser example
//
//	@Summary		Register a new User
//	@Description	Register a new User with a Username, Email and Password
//	@ID				register-a-user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.RegisterAnUserBody	true	"Register a new User"
//	@Success		201		{object}	responses.Response	"User Created!"
//	@Failure		400		{object}	responses.Response	"Invalid Body"
//	@Failure		500		{object}	responses.Response	"Failure in the Database"
//	@Router			/api/users/register [post]
func RegisterAnUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

		//TODO check if the username exist or the email is being used
		//TODO Understand better the hashFunction. why cost 14?

		//Encrypt password
		password, _ := auth.HashPassword(string(user.Password))

		//Data to be sended to the server
		newUser := models.User{
			ID:       primitive.NewObjectID(),
			Username: user.Username,
			Email:    user.Email,
			Password: password,
			Pets:     make([]primitive.ObjectID, 0),
			Vivarium: make([]primitive.ObjectID, 0),
			Money:    100,
		}

		_, err := userCollections.InsertOne(ctx, newUser)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusCreated, "success", map[string]interface{}{"data": "User created!"})
	}
}

// LoginAnUser example
//
//	@Summary		Login an User
//	@Description	Login an User with a Username and Password
//	@ID				login-an-user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.UserCredentials	true	"Logs an User"
//	@Success		200		{object}	responses.Response		"Logged User and return JWT cookie"
//	@Failure		400		{object}	responses.Response		"Invalid body"
//	@Failure		404		{object}	responses.Response		"User not found"
//	@Failure		500		{object}	responses.Response		"Failure in the Database"
//	@Router			/api/users/login [post]
func LoginAnUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
			responses.EncodeResponse(rw, http.StatusNotFound, "error", map[string]interface{}{"data": err.Error()})
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
		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": foundedUser.GetUserDetails()})
	}
}

// TODO have a renew token method

// LogoutAnUser example
//
//	@Summary		Logout an User
//	@Description	Logout current logged in user by expiring his Jwt auth
//	@ID				logout-an-user
//	@Tags			User
//	@Produce		json
//	@Success		200		{object}	responses.Response		"Logout successfuly!"
//	@Failure		500		{object}	responses.Response		"Failure on the Server"
//	@Router			/api/users/logout [post]
func LogoutAnUser() http.HandlerFunc {
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

// TODO create a Struct for user profile

// GetAnUserProfile example
//
//		@Summary		Get an User Details (Profile)
//		@Description	Receives Details from a specific User (without password)
//		@ID				get-an-user-profile
//		@Tags			User
//		@Param        	username   path  	string true  "User's name"
//		@Produce		json
//		@Success		200		{object}	models.User				"User Details"
//	 	@Failure		404		{object}	responses.Response		"User not found"
//		@Failure		500		{object}	responses.Response		"Failure on the Server"
//		@Router			/api/users/{username} [get]
func GetAnUserProfile() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
			responses.EncodeResponse(rw, http.StatusNotFound, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"user": user.GetUserDetails()})
	}
}

// LinkAPetToAUser example
//
//		@Summary		Link a created Pet to an User
//		@Description	When creating a pet successfully, this functions link to the user that it created based on the JWT
//		@ID				link-a-pet-to-user
//		@Tags			User
//		@Security 		jwt
//		@Param        	petID   path   		string true  "Pet ID"
//		@Produce		json
//		@Success		200		{object}	responses.Response		"Pet linked successfuly"
//		@Failure		401		{object}	responses.Response		"This pet has other owner!"
//	 	@Failure		404		{object}	responses.Response		"User/Pet not found"
//		@Failure		500		{object}	responses.Response		"Failure on the Server"
//		@Router			/api/users/pet/{petID} [put]
func LinkAPetToAUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		issuerContext := r.Context().Value(models.HttpContextStruct{}).(models.HttpContextStruct)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
		params := mux.Vars(r)
		petID := params["petID"]
		defer cancel()

		issuer := issuerContext.JwtIssuer
		var foundedUser models.User
		var foundedPet models.Pet

		//TODO DOT NOT LET THIS UNCHECKED
		userID, _ := primitive.ObjectIDFromHex(issuer)
		err := userCollections.FindOne(ctx, bson.M{"_id": userID}).Decode(&foundedUser)
		//Check if the user from the issuer jwt exists
		if err != nil {
			if err == mongo.ErrNoDocuments {
				responses.EncodeResponse(rw, http.StatusNotFound, "error", map[string]interface{}{"data": err.Error()})
				return
			}
		}

		petIDPrimitive, _ := primitive.ObjectIDFromHex(petID)
		//Check if the pet exists
		err = petCollection.FindOne(ctx, bson.M{"_id": petIDPrimitive}).Decode(&foundedPet)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				responses.EncodeResponse(rw, http.StatusNotFound, "error", map[string]interface{}{"data": err.Error()})
				return
			}
		}
		//CHECK IF IT HAS ONE OWNER ALREADzY
		if foundedPet.OwnerID != userID {
			responses.EncodeResponse(rw, http.StatusUnauthorized, "error", map[string]interface{}{"data": "This pet has other owner!"})
			return
		}

		foundedUser.AddPet(petIDPrimitive)
		_, err = userCollections.UpdateByID(ctx, userID, bson.M{"$set": bson.M{"pets": foundedUser.Pets}})
		// println(foundedUser.Pets)
		if err != nil {
			responses.EncodeResponse(rw, http.StatusInternalServerError, "error", map[string]interface{}{"data": err.Error()})
			return
		}

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": "Linked successfuly"})
	}
}

//TODO to change
/*
func CheckAuthenticatedUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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

		responses.EncodeResponse(rw, http.StatusOK, "success", map[string]interface{}{"data": foundedUser.GetUserDetails()})

	}
}
*/

func ChangePasswordOfAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
