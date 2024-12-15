package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type ProductResponse struct {
	ID          uint64           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Active      bool             `json:"active"`
	CategoryID  uint64           `json:"categoryID"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

func NewProductResponse(product *domain.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

type ProductPaginated struct {
	Paginated
	Product []ProductResponse `json:"products"`
}

func NewProductPaginated(products []domain.Product, total int64, page int, limit int) *ProductPaginated {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, *NewProductResponse(&product))
	}

	return &ProductPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Product: productResponses,
	}
}
