package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egon89/go-zipcode-weather/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type MockGetWeatherByZipcode struct{}

func (m *MockGetWeatherByZipcode) Execute(zipcode string) (usecase.GetWeatherByZipcodeOutputDto, error) {
	return usecase.GetWeatherByZipcodeOutputDto{
		TempCelcius:    25.0,
		TempFahrenheit: 77.0,
		TempKelvin:     298.14,
	}, nil
}

func TestGetWeather(t *testing.T) {
	mockUsecase := &MockGetWeatherByZipcode{}
	handler := NewWeatherHandler(mockUsecase)

	req, _ := http.NewRequest("GET", "/weather/12345678", nil)
	res := httptest.NewRecorder()

	handler.GetWeather(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Contains(t, res.Body.String(), `"temp_C":25`)
	assert.Contains(t, res.Body.String(), `"temp_F":77`)
	assert.Contains(t, res.Body.String(), `"temp_K":298.1`)
}
