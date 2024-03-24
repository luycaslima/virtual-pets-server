package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/auth"
	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/controllers"
	"github.com/luycaslima/virtual-pets-server/usecase"
)

func UserRoutes(router *mux.Router, useCase *usecase.UserService) {
	router.HandleFunc("/users/register", controllers.RegisterAnUser(useCase)).Methods("POST")
	router.HandleFunc("/users/login", controllers.LoginAnUser(useCase)).Methods("POST")
	router.HandleFunc("/users/{username}", configs.CacheControlWrapper(controllers.GetAnUserProfile(useCase))).Methods("GET")
	router.HandleFunc("/users/pet/adopt", auth.ValidateJWTToken(controllers.AdoptAPet(useCase))).Methods("POST")
}
