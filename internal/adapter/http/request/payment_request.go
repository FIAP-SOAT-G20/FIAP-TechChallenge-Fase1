package request

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

// PaymentPathParam contains the path parameters for the payment
type PaymentPathParam struct {
	OrderID string `uri:"order_id" binding:"required"`
}

// CreatePaymentRequest contains the request to create a payment
type CreatePaymentRequest struct {
	ExternalReference string         `json:"external_reference"`
	TotalAmount       float32        `json:"total_amount"`
	Items             []ItemsRequest `json:"items"`
	Title             string         `json:"title"`
	Description       string         `json:"description"`
	NotificationURL   string         `json:"notification_url"`
}

// ItemsRequest contains the request to create a payment item
type ItemsRequest struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint64  `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float32 `json:"total_amount"`
}

// NewPaymentRequest creates a new payment request
func NewPaymentRequest(payment *domain.CreatePaymentIN) *CreatePaymentRequest {
	if payment == nil {
		return nil
	}

	items := make([]ItemsRequest, 0)
	for _, item := range payment.Items {
		items = append(items, ItemsRequest{
			Title:       item.Title,
			Description: item.Description,
			UnitPrice:   item.UnitPrice,
			Category:    item.Category,
			UnitMeasure: item.UnitMeasure,
			Quantity:    item.Quantity,
			TotalAmount: item.TotalAmount,
		})
	}

	return &CreatePaymentRequest{
		ExternalReference: payment.ExternalReference,
		TotalAmount:       payment.TotalAmount,
		Items:             items,
		Title:             payment.Title,
		Description:       payment.Description,
		NotificationURL:   payment.NotificationUrl,
	}
}

// UpdatePaymentRequest contains the request to update a payment
type UpdatePaymentRequest struct {
	Resource string `json:"resource" example:"c16896f0-483b-4573-a493-f4d2eb59ba31"`
	Topic    string `json:"topic" enum:"payment" example:"payment"`
}

// ToDomain converts UpdatePaymentRequest to domain.UpdatePaymentIN
func (r UpdatePaymentRequest) ToDomain() *domain.UpdatePaymentIN {
	return &domain.UpdatePaymentIN{
		Resource: r.Resource,
		Topic:    r.Topic,
	}
}
