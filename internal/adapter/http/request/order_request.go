package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

// CreateOrderRequest contains the request to create an order
type CreateOrderRequest struct {
	CustomerID uint64 `json:"customer_id" binding:"required" example:"1"`
}

// ToDomain converts CreateOrderRequest to domain.Order
func (r CreateOrderRequest) ToDomain() *domain.Order {
	return &domain.Order{
		CustomerID: r.CustomerID,
	}
}

// UpdateOrderRequest contains the request to update an order
type UpdateOrderRequest struct {
	StaffID *uint64            `json:"staff_id" example:"1" ommitempty:"true"`
	Status  domain.OrderStatus `json:"status" enum:"OPEN, CANCELLED, PENDING, RECEIVED, PREPARING, READY, COMPLETED" ommitempty:"true" example:"PENDING"`
}
