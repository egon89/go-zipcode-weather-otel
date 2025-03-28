package port

import (
	"context"

	"github.com/egon89/go-zipcode-weather-gateway/internal/entity"
)

type ZipcodeWeatherPortResponse struct {
	City           string  `json:"city"`
	TempCelcius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

type ZipcodeWeatherPort interface {
	GetWeatherByZipcode(ctx context.Context, zipcode string) (entity.Weather, error)
}
