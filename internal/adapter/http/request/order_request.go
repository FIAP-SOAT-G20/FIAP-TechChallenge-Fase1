package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type CreateOrderRequest struct {
	CustomerID uint64 `json:"customer_id" binding:"required" example:"1"`
}

// UpdateOrderRequest contains the request to update an order
type UpdateOrderRequest struct {
	StaffID *uint64            `json:"staff_id" example:"1" ommitempty:"true"`
	Status  domain.OrderStatus `json:"status" enum:"OPEN, CANCELLED, PENDING, RECEIVED, PREPARING, READY, COMPLETED" ommitempty:"true" example:"PENDING"`
}
