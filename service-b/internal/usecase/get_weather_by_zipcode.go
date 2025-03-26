package usecase

import (
	"log"
	"regexp"

	"github.com/egon89/go-zipcode-weather/internal/entity"
	"github.com/egon89/go-zipcode-weather/internal/errors"
	"github.com/egon89/go-zipcode-weather/internal/ports"
)

type GetWeatherByZipcode struct {
	LocationPort    ports.LocationPort
	TemperaturePort ports.TemperaturePort
}

type GetWeatherByZipcodeOutputDto struct {
	City           string
	TempCelcius    float64
	TempFahrenheit float64
	TempKelvin     float64
}

type GetWeatherByZipcodeInterface interface {
	Execute(zipcode string) (GetWeatherByZipcodeOutputDto, error)
}

func NewGetWeatherByZipcode(locationPort ports.LocationPort, TemperaturePort ports.TemperaturePort) *GetWeatherByZipcode {
	return &GetWeatherByZipcode{
		LocationPort:    locationPort,
		TemperaturePort: TemperaturePort,
	}
}

func (g *GetWeatherByZipcode) Execute(zipcode string) (GetWeatherByZipcodeOutputDto, error) {
	if err := g.validateZipcode(zipcode); err != nil {
		return GetWeatherByZipcodeOutputDto{}, err
	}

	city, err := g.LocationPort.GetCityNameByZipcode(zipcode)
	if err != nil {
		return GetWeatherByZipcodeOutputDto{}, errors.ErrZipcodeNotFound
	}

	log.Printf("getting weather for city %s\n", city)

	tempCelcius, err := g.TemperaturePort.GetTemperatureByCity(city)
	if err != nil {
		return GetWeatherByZipcodeOutputDto{}, errors.ErrTemperatureNotFound
	}

	weather := entity.NewWeather(city, tempCelcius)

	return GetWeatherByZipcodeOutputDto{
		City:           weather.GetCity(),
		TempCelcius:    weather.GetTemperature(),
		TempFahrenheit: weather.GetTemperatureInFarhenheit(),
		TempKelvin:     weather.GetTemperatureInKelvin(),
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
