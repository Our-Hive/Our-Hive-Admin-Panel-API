package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	HUGGING_FACE_TOKEN string
	FB_STORAGE_BUCKET  string
	SECRET             string
	FB_ADMIN_SDK_PATH  string
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HUGGING_FACE_TOKEN = os.Getenv("HUGGING_FACE_TOKEN")
	FB_STORAGE_BUCKET = os.Getenv("FB_STORAGE_BUCKET")
	SECRET = os.Getenv("SECRET")
	FB_ADMIN_SDK_PATH = os.Getenv("FB_ADMIN_SDK_PATH")
}
