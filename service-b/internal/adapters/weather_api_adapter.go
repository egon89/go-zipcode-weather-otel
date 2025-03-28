package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/egon89/go-zipcode-weather/internal/config"
	"github.com/egon89/go-zipcode-weather/internal/utils"
	"go.opentelemetry.io/otel/attribute"
)

type WeatherApiAdapter struct{}

type weatherApiResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeatherApiAdapter() *WeatherApiAdapter {
	return &WeatherApiAdapter{}
}

func (wa *WeatherApiAdapter) GetTemperatureByCity(ctx context.Context, city string) (float64, error) {
	_, adapterSpan := utils.StartSpan(ctx)
	defer adapterSpan.End()
	adapterSpan.SetAttributes(attribute.String("city", city))

	adapterSpan.AddEvent("getting temperature by city adapter started")
	defer adapterSpan.AddEvent("getting temperature by city adapter finished")

	escapedCity := url.QueryEscape(city)
	url := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s", config.WeatherAPIBaseURL, config.WeatherAPIKey, escapedCity)

	log.Printf("[weatherApi] getting temperature for city %s (%s)\n", city, escapedCity)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[weatherApi] error calling weather api %s: %s\n", config.WeatherAPIBaseURL, err)
		return 0, err
	}
	defer resp.Body.Close()

	var response weatherApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("[weatherApi] error decoding response: %s\n", err)
		return 0, err
	}

	return response.Current.TempC, nil
}
