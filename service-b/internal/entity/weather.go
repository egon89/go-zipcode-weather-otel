package entity

import "github.com/egon89/go-zipcode-weather/internal/utils"

type TemperatureUnit string

type Weather struct {
	city        string
	temperature float64
}

func NewWeather(city string, tempCelcius float64) *Weather {
	return &Weather{
		city:        city,
		temperature: tempCelcius,
	}
}

func (w *Weather) GetCity() string {
	return w.city
}

func (w *Weather) GetTemperature() float64 {
	return w.temperature
}

func (w *Weather) GetTemperatureInFarhenheit() float64 {
	return utils.CelsiusToFahrenheit(w.temperature)
}

func (w *Weather) GetTemperatureInKelvin() float64 {
	return utils.CelsiusToKelvin(w.temperature)
}
