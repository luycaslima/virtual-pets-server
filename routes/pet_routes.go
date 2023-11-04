package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/controllers"
)

func PetRoutes(router *mux.Router) {
	router.HandleFunc("/api/pet/", controllers.CreateAPetToAUser()).Methods("POST")

	//Insert PetRoutes here
}
