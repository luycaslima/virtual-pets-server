package routes

import (
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/controllers"
	"github.com/luycaslima/virtual-pets-server/usecase"
)

func SpecieRoutes(router *mux.Router, useCase *usecase.SpecieService) {
	//TODO Most of this routes need Adm auth
	router.HandleFunc("/species", controllers.CreateASpecie(useCase)).Methods("POST")
	router.HandleFunc("/species", configs.CacheControlWrapper(controllers.GetListOfSpecies(useCase))).Methods("GET")
	// router.HandleFunc("/species/baby", configs.CacheControlWrapper(controllers.GetListOfInitialPhaseSpecies())).Methods("GET")
	router.HandleFunc("/species/{id}", configs.CacheControlWrapper(controllers.GetASpecieDetails(useCase))).Methods("GET")
	// router.HandleFunc("/species/{id}", controllers.UpdateASpecie()).Methods("PUT")
}
