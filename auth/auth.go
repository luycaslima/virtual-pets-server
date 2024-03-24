package auth

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luycaslima/virtual-pets-server/models"
	res "github.com/luycaslima/virtual-pets-server/responses"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWTToken(userID string) (string, time.Time, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiration := time.Now().Add(time.Hour * 2)
	//when the token expires
	claims["exp"] = expiration.Unix()
	//set the issuer of the token
	claims["iss"] = userID

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenStr, expiration, err
}

// TODO passing the issuer as a context is safe?
// This is passed in the routes with the function that need jwt validation as parameter
func ValidateJWTToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		//Check if there is a cookie
		if err != nil {
			if err == http.ErrNoCookie {
				res.EncodeResponse(w, http.StatusUnauthorized, "User Not Logged", false, err)
				return
			}
			res.EncodeResponse(w, http.StatusBadRequest, "Error on the server", false, err)
		}

		tknStr := cookie.Value

		token, err := jwt.ParseWithClaims(tknStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		//Check if it is valid
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				res.EncodeResponse(w, http.StatusUnauthorized, "User not Authorized", false, err)
				return
			}
			res.EncodeResponse(w, http.StatusBadRequest, "Error on the Server", false, err)
		}

		if token.Valid {
			//Get the user ID(issuer) by the jwt token
			claims := token.Claims.(jwt.MapClaims) //assert
			issuer, _ := claims.GetIssuer()        //TODO treat the error

			htttContext := models.HttpContextStruct{}
			r = r.WithContext(context.WithValue(
				r.Context(),
				htttContext,
				models.HttpContextStruct{
					JwtIssuer: issuer,
				},
			))
			next.ServeHTTP(w, r)
		}
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
