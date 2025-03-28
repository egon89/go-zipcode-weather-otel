package ports

import "context"

type LocationPort interface {
	GetCityNameByZipcode(ctx context.Context, zipcode string) (string, error)
}
