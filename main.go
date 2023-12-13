package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/luycaslima/virtual-pets-server/configs"
	_ "github.com/luycaslima/virtual-pets-server/docs"
	"github.com/luycaslima/virtual-pets-server/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title			Virtual Pets
// @version		1.0
// @description	This is the API for setting the REST functions of the Virtual Pets
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	lucasl22l@proton.me
// @host			localhost:8080
// @BasePath		/
func main() {
	//TODO this is  to allow any origin (just for the moment)
	//TODO STUDY CORS
	//https://stackoverflow.com/questions/40985920/making-golang-gorilla-cors-handler-work
	//https://dev.to/evillord666/auto-cors-preflight-handle-wih-gorillamux-and-go-855
	//corsObj := handlers.AllowedOrigins([]string{"*"})
	credentails := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})

	router := mux.NewRouter()

	//run database
	fmt.Println("Connecting Database")
	configs.ConnectDB()

	//Initialize routes
	routes.SpeciesRoutes(router)
	routes.PetRoutes(router)
	routes.UserRoutes(router)

	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	/*handlers.AllowCredentials()*/
	//TODO SETUP CORS FOR OTHER DOMAINS
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(credentails, origins)(router)))
}
