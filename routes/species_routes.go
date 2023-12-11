package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/controllers"
)

func SpeciesRoutes(router *mux.Router) {
	//TODO Create muss be a auth by Admin user
	router.HandleFunc("/api/species", controllers.CreateASpecie()).Methods("POST")
	router.HandleFunc("/api/species", controllers.GetBabyFormSpecies()).Methods("GET")
}
