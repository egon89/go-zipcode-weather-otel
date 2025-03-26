package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-zipcode-weather-gateway/internal/config"
	"github.com/egon89/go-zipcode-weather-gateway/internal/entity"
	"github.com/egon89/go-zipcode-weather-gateway/internal/errors"
	"github.com/egon89/go-zipcode-weather-gateway/internal/port"
)

type ZipcodeWeatherAdapter struct{}

func NewZipcodeWeatherAdapter() *ZipcodeWeatherAdapter {
	return &ZipcodeWeatherAdapter{}
}

func (z *ZipcodeWeatherAdapter) GetWeatherByZipcode(zipcode string) (entity.Weather, error) {
	log.Printf("[zipcode-weather] getting weather for zipcode %s\n", zipcode)

	url := fmt.Sprintf("%s/weather/%s", config.ZipcodeWeatherBaseURL, zipcode)

	resp, errReq := http.Get(url)
	if errReq != nil {
		log.Printf("[zipcode-weather] error calling zipcode-weather api %s: %s\n", config.ZipcodeWeatherBaseURL, errReq)
		return entity.Weather{}, errReq
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[zipcode-weather] error calling zipcode-weather api %s: %s\n", config.ZipcodeWeatherBaseURL, resp.Status)
		switch resp.StatusCode {
		case http.StatusNotFound:
			return entity.Weather{}, errors.ErrZipcodeNotFound
		default:
			return entity.Weather{}, fmt.Errorf("error calling zipcode-weather api %s: %s", config.ZipcodeWeatherBaseURL, resp.Status)
		}
	}

	var response port.ZipcodeWeatherPortResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("[zipcode-weather] error decoding response: %s\n", err)
		return entity.Weather{}, err
	}

	return entity.NewWeather(
		response.City, response.TempCelcius, response.TempFahrenheit, response.TempKelvin), nil
}
