package domain

import "errors"

var (
	ErrConflict           = errors.New("data conflicts with existing data")
	ErrNotFound           = errors.New("data not found")
	ErrInvalidParam       = errors.New("invalid parameter")
	ErrInvalidQueryParams = errors.New("invalid query parameters")
)

var (
	ErrTokenDuration = errors.New("invalid token duration format")
	ErrTokenCreation = errors.New("error creating token")
	ErrExpiredToken  = errors.New("access token has expired")
	ErrInvalidToken  = errors.New("access token is invalid")
)

var (
	ErrOrderInvalidStatusTransition = errors.New("invalid status transition")
	ErrOrderWithoutProducts         = errors.New("order without products")
)
