package domain

import "time"

type PaymentStatus string

const (
	PROCESSING PaymentStatus = "PROCESSING"
	CONFIRMED  PaymentStatus = "CONFIRMED"
	FAILED     PaymentStatus = "FAILED"
	CANCELED   PaymentStatus = "CANCELED"
)

type Payment struct {
	ID                uint64
	Status            PaymentStatus
	ExternalPaymentID string
	QrData            string
	OrderID           uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
