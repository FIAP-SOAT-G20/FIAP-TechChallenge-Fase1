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

type CategoriesPaginated struct {
	Paginated
	Categories []CategoryResponse `json:"categories"`
}

func NewCategoriesPaginated(categories []domain.Category, total int64, page int, limit int) *CategoriesPaginated {
	categoryResponses := make([]CategoryResponse, 0, len(categories))
	for _, customer := range categories {
		categoryResponse := NewCategoryResponse(&customer)
		if categoryResponse != nil {
			categoryResponses = append(categoryResponses, *categoryResponse)
		}
	}

	return &CategoriesPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Categories: categoryResponses,
	}
}
