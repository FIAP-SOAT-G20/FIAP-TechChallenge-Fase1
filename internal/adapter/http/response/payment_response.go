package response

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

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
