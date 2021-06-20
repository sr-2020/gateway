package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrMultiLogin = errors.New("multi login")
)
