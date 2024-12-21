package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CategoryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategoryResponse(category *domain.Category) *CategoryResponse {
	if category == nil {
		return nil
	}

	return &CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}
