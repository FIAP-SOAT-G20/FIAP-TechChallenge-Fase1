package response

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CreatePaymentResponse struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QrData         string `json:"qr_data"`
}

type PaymentResponse struct {
	ID                uint64               `json:"id"`
	Status            domain.PaymentStatus `json:"status"`
	OrderID           uint64               `json:"order_id"`
	ExternalPaymentID string               `json:"external_payment_id"`
	QrData            string               `json:"qr_data"`
}

func NewPaymentResponse(payment *domain.Payment) *PaymentResponse {
	if payment == nil {
		return nil
	}

	return &PaymentResponse{
		ID:                payment.ID,
		Status:            payment.Status,
		OrderID:           payment.OrderID,
		ExternalPaymentID: payment.ExternalPaymentID,
		QrData:            payment.QrData,
	}
}

// ToCreatePaymentOUTDomain creates a new payment request output
func ToCreatePaymentOUTDomain(payment *CreatePaymentResponse) *domain.CreatePaymentOUT {
	return &domain.CreatePaymentOUT{
		InStoreOrderID: payment.InStoreOrderID,
		QrData:         payment.QrData,
	}
}
