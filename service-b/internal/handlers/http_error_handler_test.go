package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egon89/go-zipcode-weather/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandlerHttpError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid zipcode error",
			err:            errors.ErrInvalidZipcode,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   "Invalid zipcode\n",
		},
		{
			name:           "zipcode not found error",
			err:            errors.ErrZipcodeNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "City not found\n",
		},
		{
			name:           "temperature not found error",
			err:            errors.ErrTemperatureNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Temperature not found\n",
		},
		{
			name:           "internal server error",
			err:            fmt.Errorf("some error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal server error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			HandlerHttpError(rr, tt.err)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
