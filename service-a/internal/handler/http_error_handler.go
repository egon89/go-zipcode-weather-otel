package handler

import (
	"net/http"

	"github.com/egon89/go-zipcode-weather-gateway/internal/errors"
)

func HandlerHttpError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch err {
	case errors.ErrInvalidZipcode:
		http.Error(w, "Invalid zipcode", http.StatusUnprocessableEntity)
	case errors.ErrZipcodeNotFound:
		http.Error(w, "City not found", http.StatusNotFound)
	default:
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
