package utils

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading .env file")
	}

	return os.Getenv("MONGOURI")
}