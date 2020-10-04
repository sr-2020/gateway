package domain

import "errors"

var (
	ErrNotFound = errors.New("Not found")
	ErrBadRequest = errors.New("Bad request")
	ErrUnauthorized = errors.New("Unauthorized")
)
