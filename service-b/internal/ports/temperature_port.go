package ports

import "context"

type TemperaturePort interface {
	GetTemperatureByCity(ctx context.Context, city string) (float64, error)
}
