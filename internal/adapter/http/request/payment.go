package request

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CreatePaymentRequest struct {
	ExternalReference string  `json:"external_reference"`
	TotalAmount       float32 `json:"total_amount"`
	Items             []Items `json:"items"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	NotificationURL   string  `json:"notification_url"`
}

type Items struct {
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    uint64  `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float32 `json:"total_amount"`
}

func NewPaymentRequest(payment *domain.CreatePaymentIN) *CreatePaymentRequest {
	if payment == nil {
		return nil
	}

	items := make([]Items, 0)
	for _, item := range payment.Items {
		items = append(items, Items{
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
		ExternalReference: payment.OrderID,
		TotalAmount:       payment.TotalAmount,
		Items:             items,
		Title:             payment.Title,
		Description:       payment.Description,
		NotificationURL:   payment.Webhook,
	}
}
