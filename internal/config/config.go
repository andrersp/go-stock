package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	STAGE    string
	API_PORT string
)

func LoadConfig() {
	STAGE = os.Getenv("STAGE")
	err := godotenv.Load()

	if err != nil && STAGE == "LOCAL" {
		log.Fatal(err)
	}
	API_PORT = os.Getenv("API_PORT")

}
