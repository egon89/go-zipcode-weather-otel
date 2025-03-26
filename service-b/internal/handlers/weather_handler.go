package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-zipcode-weather/internal/usecase"
	"github.com/egon89/go-zipcode-weather/internal/utils"
	"github.com/go-chi/chi/v5"
)

type WeatherHandler struct {
	usecase usecase.GetWeatherByZipcodeInterface
}

type GetWeatherResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeatherHandler(getWeatherByZipcode usecase.GetWeatherByZipcodeInterface) *WeatherHandler {
	return &WeatherHandler{
		usecase: getWeatherByZipcode,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcodeStr := chi.URLParam(r, "zipcode")

	output, err := h.usecase.Execute(zipcodeStr)
	if err != nil {
		HandlerHttpError(w, err)
		return
	}

	response := GetWeatherResponse{
		City:  output.City,
		TempC: utils.RoundToOneDecimal(output.TempCelcius),
		TempF: utils.RoundToOneDecimal(output.TempFahrenheit),
		TempK: utils.RoundToOneDecimal(output.TempKelvin),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
