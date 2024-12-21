package response

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CustomerResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func NewCustomerResponse(customer *domain.Customer) *CustomerResponse {
	if customer == nil {
		return nil
	}

	return &CustomerResponse{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: customer.Email,
		CPF:   customer.CPF,
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
