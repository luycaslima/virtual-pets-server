package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/controllers"
)

func UserRoutes(router *mux.Router) {

	router.HandleFunc("/api/users", controllers.CreateAUser()).Methods("POST")
	router.HandleFunc("/api/users/{username}", controllers.GetAUser()).Methods("GET")
}
