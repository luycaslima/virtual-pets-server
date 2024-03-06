package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/auth"
	"github.com/luycaslima/virtual-pets-server/controllers"
)

func UserRoutes(router *mux.Router) {

	router.HandleFunc("/api/users/register", controllers.RegisterAnUser()).Methods("POST")
	router.HandleFunc("/api/users/login", controllers.LoginAnUser()).Methods("POST")
	router.HandleFunc("/api/users/logout", controllers.LogoutAnUser()).Methods("POST")
	router.HandleFunc("/api/users/{username}", controllers.GetAnUserProfile()).Methods("GET")
	//router.HandleFunc("/api/users", controllers.CheckAuthenticatedUser()).Methods("GET")
	router.HandleFunc("/api/users/pet/{petID}", auth.ValidateJWT(controllers.LinkAPetToAUser())).Methods("PUT")
}
