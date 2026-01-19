package usecase

import "errors"

var (
	errInvalidVaue error = errors.New("error: invalid value for")
	errOccurred    error = errors.New("an error has occurred")
	errNotFound    error = errors.New("error: object not found")
)
