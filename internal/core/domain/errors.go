package domain

import "errors"

var (
	ErrConflict = errors.New("data conflicts with existing data")
	ErrNotFound = errors.New("data not found")
)

var (
	ErrTokenDuration = errors.New("invalid token duration format")
	ErrTokenCreation = errors.New("error creating token")
	ErrExpiredToken  = errors.New("access token has expired")
	ErrInvalidToken  = errors.New("access token is invalid")
)
