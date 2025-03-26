package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port                  string
	ZipcodeWeatherBaseURL string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	Port = GetEnv("PORT", "8080")
	ZipcodeWeatherBaseURL = GetEnv("ZIPCODE_WEATHER_BASE_URL", "")
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
