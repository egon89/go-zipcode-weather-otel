package errors

import "fmt"

var (
	ErrInvalidZipcode  = fmt.Errorf("invalid zipcode")
	ErrZipcodeNotFound = fmt.Errorf("can not find zipcode")
)
