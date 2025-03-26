package ports

type TemperaturePort interface {
	GetTemperatureByCity(city string) (float64, error)
}
