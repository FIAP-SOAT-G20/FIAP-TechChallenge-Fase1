package domain

import "errors"

var (
	ErrConflict = errors.New("data conflicts with existing data")
	ErrNotFound = errors.New("data not found")
)
