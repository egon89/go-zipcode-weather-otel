package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/egon89/go-zipcode-weather-gateway/internal/config"
	"github.com/egon89/go-zipcode-weather-gateway/internal/entity"
	"github.com/egon89/go-zipcode-weather-gateway/internal/errors"
	"github.com/egon89/go-zipcode-weather-gateway/internal/port"
	"github.com/egon89/go-zipcode-weather-gateway/internal/util"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

type ZipcodeWeatherAdapter struct{}

func NewZipcodeWeatherAdapter() *ZipcodeWeatherAdapter {
	return &ZipcodeWeatherAdapter{}
}

func (z *ZipcodeWeatherAdapter) GetWeatherByZipcode(ctx context.Context, zipcode string) (entity.Weather, error) {
	log.Printf("[zipcode-weather] getting weather for zipcode %s\n", zipcode)
	adapterCtx, adapterSpan := util.StartSpan(ctx)
	adapterSpan.SetAttributes(attribute.String("zipcode", zipcode))

	url := fmt.Sprintf("%s/weather/%s", config.ZipcodeWeatherBaseURL, zipcode)

	httpClient := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	request, err := http.NewRequestWithContext(adapterCtx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("error creating request to zipcode-weather service: %v", err)
		adapterSpan.End()
		return entity.Weather{}, err
	}

	resp, err := httpClient.Do(request)
	adapterSpan.End()

	if err != nil {
		log.Printf("error calling zipcode-weather service: %v", err)
		return entity.Weather{}, err
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
