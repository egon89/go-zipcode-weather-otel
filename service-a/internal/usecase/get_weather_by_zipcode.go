package usecase

import (
	"context"
	"log"
	"regexp"

	"github.com/egon89/go-zipcode-weather-gateway/internal/errors"
	"github.com/egon89/go-zipcode-weather-gateway/internal/port"
)

type GetWeatherByZipcode struct {
	ZipcodeWeatherPort port.ZipcodeWeatherPort
}

type GetWeatherByZipcodeOutputDto struct {
	City           string
	TempCelcius    float64
	TempFahrenheit float64
	TempKelvin     float64
}

type GetWeatherByZipcodeInterface interface {
	Execute(ctx context.Context, zipcode string) (GetWeatherByZipcodeOutputDto, error)
}

func NewGetWeatherByZipcode(zipcodeWeatherPort port.ZipcodeWeatherPort) *GetWeatherByZipcode {
	return &GetWeatherByZipcode{
		ZipcodeWeatherPort: zipcodeWeatherPort,
	}
}

func (g *GetWeatherByZipcode) Execute(ctx context.Context, zipcode string) (GetWeatherByZipcodeOutputDto, error) {
	if err := g.validateZipcode(zipcode); err != nil {
		return GetWeatherByZipcodeOutputDto{}, err
	}

	weather, err := g.ZipcodeWeatherPort.GetWeatherByZipcode(ctx, zipcode)
	if err != nil {
		log.Printf("error getting weather by zipcode %s: %v\n", zipcode, err)
		return GetWeatherByZipcodeOutputDto{}, err
	}

	return GetWeatherByZipcodeOutputDto{
		City:           weather.GetCity(),
		TempCelcius:    weather.GetTemperatureCelcius(),
		TempFahrenheit: weather.GetTemperatureFahrenheit(),
		TempKelvin:     weather.GetTemperatureKelvin(),
	}, nil
}

func (g *GetWeatherByZipcode) validateZipcode(zipcode string) error {
	regex := `^\d{8}$`
	matched, err := regexp.MatchString(regex, zipcode)
	if err != nil || !matched {
		return errors.ErrInvalidZipcode
	}

	return nil
}
