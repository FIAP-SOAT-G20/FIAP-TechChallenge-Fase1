package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type ProductResponse struct {
	ID          uint64            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Active      bool              `json:"active"`
	CategoryID  uint64            `json:"category_id"`
	Category    *CategoryResponse `json:"category,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func NewProductResponse(product *domain.Product) ProductResponse {
	if product == nil {
		return ProductResponse{}
	}

	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		Category:    NewCategoryResponse(&product.Category),
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

type ProductPaginated struct {
	Paginated
	Products []ProductResponse `json:"products"`
}

func NewProductPaginated(products []domain.Product, total int64, page int, limit int) ProductPaginated {
	productResponses := make([]ProductResponse, 0, len(products))
	for _, product := range products {
		productResponses = append(productResponses, NewProductResponse(&product))
	}

	return ProductPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Products: productResponses,
	}
}
