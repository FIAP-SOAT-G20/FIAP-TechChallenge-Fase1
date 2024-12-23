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

type CreatePayment struct {
	ExternalReference string
	TotalAmount       float32
	Items             []Items
	Title             string
	Description       string
	NotificationURL   string
}

type Items struct {
	Category    string
	Title       string
	Description string
	UnitPrice   float32
	Quantity    uint64
	UnitMeasure string
	TotalAmount float32
}

type CreatePaymentResponse struct {
	InStoreOrderID string
	QrData         string
}
