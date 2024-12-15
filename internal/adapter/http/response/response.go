package response

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type Paginated struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type ProductPaginated struct {
	Paginated
	Products []domain.Product `json:"products"`
}

type CustomersPaginated struct {
	Paginated
	Customers []domain.Customer `json:"customers"`
}
