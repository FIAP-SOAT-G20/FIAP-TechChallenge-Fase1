package response

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CustomerResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CPF       string    `json:"cpf"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCustomerResponse(customer *domain.Customer) *CustomerResponse {
	if customer == nil {
		return nil
	}

	return &CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		CPF:       customer.CPF,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}

type CustomersPaginated struct {
	Paginated
	Customers []CustomerResponse `json:"customers"`
}

func NewCustomersPaginated(customers []domain.Customer, total int64, page int, limit int) *CustomersPaginated {
	customerResponses := make([]CustomerResponse, 0, len(customers))
	for _, customer := range customers {
		customerResponse := NewCustomerResponse(&customer)
		if customerResponse != nil {
			customerResponses = append(customerResponses, *customerResponse)
		}
	}

	return &CustomersPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Customers: customerResponses,
	}
}
