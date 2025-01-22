package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"BK Mega Stacker 2.0"`
	Description string  `json:"description" binding:"required" example:"The best burger in the world"`
	Price       float32 `json:"price" binding:"required" example:"29.90"`
	CategoryID  uint64  `json:"category_id" binding:"required" example:"1"`
	Active      bool    `json:"active" example:"true" default:"true"`
}

// ToDomain converts CreateProductRequest to domain.Product
func (r CreateProductRequest) ToDomain() *domain.Product {
	return &domain.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		CategoryID:  r.CategoryID,
		Active:      r.Active,
	}
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name        string  `json:"name" example:"McDonald's Big Mac"`
	Description string  `json:"description" example:"The best burger in the world"`
	Price       float32 `json:"price" example:"29.90"`
	CategoryID  uint64  `json:"category_id" example:"1"`
	Active      bool    `json:"active" example:"true"`
}

// ToDomain converts UpdateProductRequest to domain.Product
func (r UpdateProductRequest) ToDomain(id uint64) *domain.Product {
	return &domain.Product{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		CategoryID:  r.CategoryID,
		Active:      r.Active,
	}
}
