package main

import (
	"fmt"

	"github.com/luycaslima/virtual-pets-server/database"
	"github.com/luycaslima/virtual-pets-server/server"
)

func main() {
	db := database.ConnectMongoDB()

	muxServer := server.NewMuxServer(db.GetDb())

	err := muxServer.StartServer(":8080")
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
