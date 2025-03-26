package handler

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/go-zipcode-weather-gateway/internal/usecase"
)

type WeatherHandler struct {
	usecase usecase.GetWeatherByZipcodeInterface
}

type WeatherRequest struct {
	Zipcode string `json:"cep"`
}

type WeatherResponse struct {
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

func (h *WeatherHandler) WeatherByZipcode(w http.ResponseWriter, r *http.Request) {
	var req WeatherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	output, err := h.usecase.Execute(req.Zipcode)
	if err != nil {
		HandlerHttpError(w, err)
		return
	}

	response := WeatherResponse{
		City:  output.City,
		TempC: output.TempCelcius,
		TempF: output.TempFahrenheit,
		TempK: output.TempKelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
