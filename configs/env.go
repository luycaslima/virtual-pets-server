package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load the enviroment file variable!")
	}
}
