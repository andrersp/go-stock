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
	JWT_ACCESS_TOKEN_SECRET, JWT_REFRESH_TOKEN_SECRET              string
)

func LoadConfig() {

	err := godotenv.Load()
	STAGE = os.Getenv("STAGE")

	if err != nil && STAGE == DEV_STAGE {
		log.Fatal(err)
	}
	API_PORT = os.Getenv("API_PORT")
	JWT_REFRESH_TOKEN_SECRET = os.Getenv("JWT_REFRESH_TOKEN_SECRET")
	JWT_ACCESS_TOKEN_SECRET = os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWD = os.Getenv("DB_PASSWD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_REQUIRED_SSL = os.Getenv("DB_REQUIRED_SSL")

}
