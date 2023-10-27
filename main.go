package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luycaslima/virtual-pets-server/configs"
)

func main() {
	router := mux.NewRouter()

	//run database
	configs.ConnectDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
