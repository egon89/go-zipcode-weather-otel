package adapters

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egon89/go-zipcode-weather/internal/config"
	"github.com/stretchr/testify/assert"
)

func mockViaCepResponse(cep string) string {
	return fmt.Sprintf(`{
        "cep": "%s",
        "logradouro": "Main Street",
        "complemento": "Apartment 123",
        "bairro": "Downtown",
        "localidade": "New York",
        "uf": "NY"
    }`, cep)
}

func TestViaCepAdapter_GetCityNameByZipcode_Success(t *testing.T) {
	cep := "01153000"
	expectedCity := "New York"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockViaCepResponse(cep)))
	}))
	defer ts.Close()

	originalURL := config.ViaCepBaseURL
	config.ViaCepBaseURL = ts.URL
	defer func() {
		config.ViaCepBaseURL = originalURL
	}()

	adapter := NewViaCepAdapter()
	city, err := adapter.GetCityNameByZipcode(context.Background(), cep)

	assert.NoError(t, err)
	assert.Equal(t, expectedCity, city)
}

func TestViaCepAdapter_GetCityNameByZipcode_NotFound(t *testing.T) {
	cep := "01153000"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"localidade": ""}`))
	}))
	defer ts.Close()

	originalURL := config.ViaCepBaseURL
	config.ViaCepBaseURL = ts.URL
	defer func() {
		config.ViaCepBaseURL = originalURL
	}()

	adapter := NewViaCepAdapter()
	city, err := adapter.GetCityNameByZipcode(context.Background(), cep)

	assert.Error(t, err)
	assert.Equal(t, "", city)
	assert.Contains(t, err.Error(), fmt.Sprintf("city not found for zipcode %s", cep))
}
