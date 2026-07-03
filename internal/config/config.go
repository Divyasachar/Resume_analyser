package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppVersion   string
	AppPort      string
	UploadFolder string
	OutputFolder string
	LogLevel     string
}

var AppConfig Config

func Load() error {

	err := godotenv.Load("configs/app.env")
	if err != nil {
		log.Println("No app.env found, using environment variables")
	}

	AppConfig = Config{
		AppName:      getEnv("APP_NAME", "Resume Analyzer"),
		AppVersion:   getEnv("APP_VERSION", "1.0.0"),
		AppPort:      getEnv("APP_PORT", "8080"),
		UploadFolder: getEnv("UPLOAD_FOLDER", "uploads"),
		OutputFolder: getEnv("OUTPUT_FOLDER", "outputs"),
		LogLevel:     getEnv("LOG_LEVEL", "INFO"),
	}

	return nil
}

func getEnv(key, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}