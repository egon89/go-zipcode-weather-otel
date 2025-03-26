package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
		{37, 98.60000000000001},
	}

	for _, test := range tests {
		result := CelsiusToFahrenheit(test.input)

		assert.Equal(t, test.expected, result)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{0, 273.15},
		{100, 373.15},
		{-273.15, 0},
		{25, 298.15},
	}

	for _, test := range tests {
		result := CelsiusToKelvin(test.input)

		assert.Equal(t, test.expected, result)
	}
}

func TestRoundToOneDecimal(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{2.345, 2.3},
		{2.355, 2.4},
		{2.0, 2.0},
		{2.05, 2.1},
	}

	for _, test := range tests {
		result := RoundToOneDecimal(test.input)

		assert.Equal(t, test.expected, result)
	}
}
