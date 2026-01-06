package postgres

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrUnsupported  = errors.New("unsupported filter key")
)
