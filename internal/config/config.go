package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const DEV_STAGE = "LOCAL"

var (
	STAGE                                                          string
	API_PORT                                                       string
	DB_HOST, DB_USER, DB_PASSWD, DB_NAME, DB_PORT, DB_REQUIRED_SSL string
)

func LoadConfig() {

	err := godotenv.Load()
	STAGE = os.Getenv("STAGE")

	if err != nil && STAGE == DEV_STAGE {
		log.Fatal(err)
	}
	API_PORT = os.Getenv("API_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWD = os.Getenv("DB_PASSWD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_REQUIRED_SSL = os.Getenv("DB_REQUIRED_SSL")

}
