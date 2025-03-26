package usecase

import (
	"testing"

	"github.com/egon89/go-zipcode-weather/internal/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocationPort struct {
	mock.Mock
}

func (m *MockLocationPort) GetCityNameByZipcode(zipcode string) (string, error) {
	args := m.Called(zipcode)
	return args.String(0), args.Error(1)
}

type MockTemperaturePort struct {
	mock.Mock
}

func (m *MockTemperaturePort) GetTemperatureByCity(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

func TestGetWeatherByZipcode_Execute(t *testing.T) {
	t.Run("invalid zipcode", func(t *testing.T) {
		mockLocationPort := new(MockLocationPort)
		mockTemperaturePort := new(MockTemperaturePort)
		usecase := NewGetWeatherByZipcode(mockLocationPort, mockTemperaturePort)

		_, err := usecase.Execute("123")

		assert.Equal(t, errors.ErrInvalidZipcode, err)
	})

	t.Run("zipcode not found", func(t *testing.T) {
		mockLocationPort := new(MockLocationPort)
		mockTemperaturePort := new(MockTemperaturePort)
		usecase := NewGetWeatherByZipcode(mockLocationPort, mockTemperaturePort)
		mockLocationPort.On("GetCityNameByZipcode", "12345678").Return("", errors.ErrZipcodeNotFound)

		_, err := usecase.Execute("12345678")

		assert.Equal(t, errors.ErrZipcodeNotFound, err)
		mockLocationPort.AssertExpectations(t)
	})

	t.Run("temperature not found", func(t *testing.T) {
		mockLocationPort := new(MockLocationPort)
		mockTemperaturePort := new(MockTemperaturePort)
		usecase := NewGetWeatherByZipcode(mockLocationPort, mockTemperaturePort)
		mockLocationPort.On("GetCityNameByZipcode", "12345678").Return("TestCity", nil)
		mockTemperaturePort.On("GetTemperatureByCity", "TestCity").Return(0.0, errors.ErrTemperatureNotFound)

		_, err := usecase.Execute("12345678")

		assert.Equal(t, errors.ErrTemperatureNotFound, err)
		mockLocationPort.AssertExpectations(t)
		mockTemperaturePort.AssertExpectations(t)
	})

	t.Run("successful execution", func(t *testing.T) {
		mockLocationPort := new(MockLocationPort)
		mockTemperaturePort := new(MockTemperaturePort)
		usecase := NewGetWeatherByZipcode(mockLocationPort, mockTemperaturePort)
		mockLocationPort.On("GetCityNameByZipcode", "12345678").Return("TestCity", nil)
		mockTemperaturePort.On("GetTemperatureByCity", "TestCity").Return(25.0, nil)

		result, err := usecase.Execute("12345678")

		assert.NoError(t, err)
		assert.Equal(t, 25.0, result.TempCelcius)
		assert.Equal(t, 77.0, result.TempFahrenheit)
		assert.Equal(t, 298.15, result.TempKelvin)
		mockLocationPort.AssertExpectations(t)
		mockTemperaturePort.AssertExpectations(t)
	})
}
