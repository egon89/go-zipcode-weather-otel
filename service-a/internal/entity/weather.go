package entity

type Weather struct {
	city           string
	tempCelcius    float64
	tempFahrenheit float64
	tempKelvin     float64
}

func NewWeather(city string, tempCelcius float64, tempFahrenheit float64, tempKelvin float64) Weather {
	return Weather{
		city:           city,
		tempCelcius:    tempCelcius,
		tempFahrenheit: tempFahrenheit,
		tempKelvin:     tempKelvin,
	}
}

func (w Weather) GetCity() string {
	return w.city
}

func (w Weather) GetTemperatureCelcius() float64 {
	return w.tempCelcius
}

func (w Weather) GetTemperatureFahrenheit() float64 {
	return w.tempFahrenheit
}

func (w Weather) GetTemperatureKelvin() float64 {
	return w.tempKelvin
}
