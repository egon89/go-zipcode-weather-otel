package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port              string
	ViaCepBaseURL     string
	WeatherAPIBaseURL string
	WeatherAPIKey     string
	OtelCollectorHost string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	Port = GetEnv("PORT", "8080")
	ViaCepBaseURL = GetEnv("VIA_CEP_BASE_URL", "")
	WeatherAPIBaseURL = GetEnv("WEATHER_API_BASE_URL", "")
	WeatherAPIKey = GetEnv("WEATHER_API_KEY", "")
	OtelCollectorHost = GetEnv("OTEL_COLLECTOR_HOST", "")
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
