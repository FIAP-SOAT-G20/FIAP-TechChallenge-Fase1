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

type CreatePaymentIN struct {
	ExternalReference string
	TotalAmount       float32
	Items             []ItemsIN
	Title             string
	Description       string
	NotificationUrl   string
}

type ItemsIN struct {
	Category    string
	Title       string
	Description string
	UnitPrice   float32
	Quantity    uint64
	UnitMeasure string
	TotalAmount float32
}

type CreatePaymentOUT struct {
	OrderID string
	QrData  string
}
