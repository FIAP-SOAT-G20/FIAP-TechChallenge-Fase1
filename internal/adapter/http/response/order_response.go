package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderResponse struct {
	ID         uint64                 `json:"id"`
	CustomerID uint64                 `json:"customer_id"`
	TotalBill  float32                `json:"total_bill"`
	Status     string                 `json:"status"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	Products   []OrderProductResponse `json:"products"`
}

func NewOrderResponse(order *domain.Order) OrderResponse {
	if order == nil {
		return OrderResponse{}
	}

	products := make([]OrderProductResponse, 0, len(order.OrderProducts))
	for _, orderProduct := range order.OrderProducts {
		products = append(products, NewOrderProductResponse(&orderProduct))
	}

	return OrderResponse{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     order.Status.ToString(),
		TotalBill:  order.TotalBill,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Products:   products,
	}
}

type OrderPaginated struct {
	Paginated
	Orders []OrderResponse `json:"orders"`
}

func NewOrderPaginated(orders []domain.Order, total int64, page int, limit int) OrderPaginated {
	responses := make([]OrderResponse, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, NewOrderResponse(&order))
	}

	return OrderPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Orders: responses,
	}
}
