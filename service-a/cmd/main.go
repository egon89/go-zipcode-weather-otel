package main

import (
	"log"
	"net/http"

	"github.com/egon89/go-zipcode-weather-gateway/internal/adapter"
	"github.com/egon89/go-zipcode-weather-gateway/internal/config"
	"github.com/egon89/go-zipcode-weather-gateway/internal/handler"
	"github.com/egon89/go-zipcode-weather-gateway/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func main() {
	config.LoadEnv()

	r := chi.NewRouter()

	zipcodeWeatherAdapter := adapter.NewZipcodeWeatherAdapter()
	weatherHandler := handler.NewWeatherHandler(
		usecase.NewGetWeatherByZipcode(zipcodeWeatherAdapter))

	r.Post("/weather", weatherHandler.WeatherByZipcode)

	log.Println("Starting server on port " + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
