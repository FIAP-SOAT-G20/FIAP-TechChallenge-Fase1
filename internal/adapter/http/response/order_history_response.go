package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderHistoryResponse struct {
	ID        uint64    `json:"id"`
	OrderID   uint64    `json:"order_id"`
	StaffID   *uint64   `json:"staff_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewOrderHistoryResponse(order *domain.OrderHistory) OrderHistoryResponse {
	if order == nil {
		return OrderHistoryResponse{}
	}

	return OrderHistoryResponse{
		ID:        order.ID,
		OrderID:   order.OrderID,
		Status:    order.Status.ToString(),
		StaffID:   order.StaffID,
		CreatedAt: order.CreatedAt,
	}
}

type OrderHistoryPaginated struct {
	Paginated
	OrderHistories []OrderHistoryResponse `json:"order_histories"`
}

func NewOrderHistoryPaginated(orders []domain.OrderHistory, total int64, page int, limit int) OrderHistoryPaginated {
	responses := make([]OrderHistoryResponse, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, NewOrderHistoryResponse(&order))
	}

	return OrderHistoryPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		OrderHistories: responses,
	}
}
