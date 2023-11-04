package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/configs"
	"github.com/luycaslima/virtual-pets-server/routes"
)

func main() {
	//TODO this is  to allow any origin (just for the moment)
	//TODO STUDY CORS
	//https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
	//https://dev.to/evillord666/auto-cors-preflight-handle-wih-gorillamux-and-go-855
	//corsObj := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter()

	//run database
	fmt.Println("Connecting Database")
	configs.ConnectDB()

	//Initialize routes
	routes.PetRoutes(router)
	routes.UserRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowCredentials())(router)))
}
