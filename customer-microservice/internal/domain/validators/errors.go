package validators

import "errors"

var (
	errEmptyFields  error = errors.New("error: empty value for fied")
	errTooLong      error = errors.New("error: too long value for fied ")
	errInvalidGenda error = errors.New("error: invalid input genda ")
	errInvalidId    error = errors.New("error: provided is invalid")
	errInvalidPhone error = errors.New("error: invalid input phone number")
)
