package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load the MONGODB Enviroment variable from the .env File
func GetEnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't  load the enviroment file variable!")
	}

	return os.Getenv("MONGOURI")
}
