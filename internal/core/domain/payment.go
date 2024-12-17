package domain

import "time"

type PaymentStatus string

const (
	PENDING   PaymentStatus = "PENDING"
	CONFIRMED PaymentStatus = "CONFIRMED"
	CANCELED  PaymentStatus = "CANCELED"
)

type Payment struct {
	ID        uint64
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
