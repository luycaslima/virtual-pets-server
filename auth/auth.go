package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWT(userID string) (string, time.Time, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiration := time.Now().Add(time.Hour)
	//when the token expires
	claims["exp"] = expiration.Unix()
	// set the issuer of the token
	claims["iss"] = userID

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenStr, expiration, err
}

// This is passed in the routes with the function that need jwt validation as parameter
// how get the issuer from the token in this form?
func ValidateJWT(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// if r.Header["Token"] != nil {
		// 	token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
		// 		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		// 		if !ok {
		// 			w.WriteHeader(http.StatusUnauthorized)
		// 			w.Write([]byte("not authorized"))
		// 		}
		// 		return os.Getenv("JWT_SECRET_KEY"), nil
		// 	})

		// 	if err != nil {
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		w.Write([]byte("not authorized " + err.Error()))
		// 	}

		// 	//pass middleware
		// 	if token.Valid {
		// 		next.ServeHTTP(w, r) // func(w http.ResponseWriter, r *http.Request)
		// 	}
		// } else {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("not authorized"))
		// }
	}
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPassword(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// func GetJWT(w http.ResponseWriter, r *http.Request) {
// 	if r.Header["Access"] != nil {
// 		if r.Header["Access"][0] == api_key {
// 			token, err := CreateJWT()
// 			if err != nil {
// 				return
// 			}
// 		}
// 	}
// }
