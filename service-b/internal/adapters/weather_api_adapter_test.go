package adapters

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egon89/go-zipcode-weather/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperatureByCity(t *testing.T) {
	tests := []struct {
		city         string
		mockResponse string
		expectedTemp float64
		expectedErr  bool
	}{
		{
			city:         "London",
			mockResponse: `{"current": {"temp_c": 15.0}}`,
			expectedTemp: 15.0,
		},
		{
			city:         "Porto Alegre",
			mockResponse: `{"current": {"temp_c": 23.5}}`,
			expectedTemp: 23.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.city, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.mockResponse))
			}))
			defer ts.Close()

			originalURL := config.WeatherAPIBaseURL
			config.WeatherAPIBaseURL = ts.URL
			defer func() {
				config.WeatherAPIBaseURL = originalURL
			}()

			wa := NewWeatherApiAdapter()
			temp, err := wa.GetTemperatureByCity(context.Background(), tt.city)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedTemp, temp)
		})
	}
}
