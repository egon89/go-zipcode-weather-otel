package ports

type LocationPort interface {
	GetCityNameByZipcode(zipcode string) (string, error)
}
