package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderProductResponse struct {
	ID        uint64    `json:"id"`
	OrderID   uint64    `json:"order_id"`
	ProductID uint64    `json:"product_id"`
	Price     float32   `json:"price"`
	Quantity  uint32    `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrderProductResponse(order *domain.OrderProduct) OrderProductResponse {
	if order == nil {
		return OrderProductResponse{}
	}

	return OrderProductResponse{
		ID:        order.ID,
		OrderID:   order.OrderID,
		ProductID: order.ProductID,
		Price:     order.Price,
		Quantity:  order.Quantity,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}

type OrderProductPaginated struct {
	Paginated
	OrderProducts []OrderProductResponse `json:"order_products"`
}

func NewOrderProductPaginated(orders []domain.OrderProduct, total int64, page int, limit int) OrderProductPaginated {
	responses := make([]OrderProductResponse, 0, len(orders))
	for _, order := range orders {
		responses = append(responses, NewOrderProductResponse(&order))
	}

	return OrderProductPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		OrderProducts: responses,
	}
}
