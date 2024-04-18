package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	HUGGING_FACE_TOKEN string
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HUGGING_FACE_TOKEN = os.Getenv("HUGGING_FACE_TOKEN")
}
