package pokeapi

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
	ErrOther    = errors.New("service error")
)
