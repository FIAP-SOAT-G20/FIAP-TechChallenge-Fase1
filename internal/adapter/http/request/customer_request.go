package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

// CreateCustomerRequest represents the request to create a customer
type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"johndoe@contact.com"`
	CPF   string `json:"cpf" binding:"required" example:"123.456.789-00"`
}

// ToDomain converts CreateCustomerRequest to domain.Customer
func (r CreateCustomerRequest) ToDomain() *domain.Customer {
	return &domain.Customer{
		Name:  r.Name,
		Email: r.Email,
		CPF:   r.CPF,
	}
}

// UpdateCustomerRequest represents the request to update a customer
type UpdateCustomerRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"johndoe@email.com"`
	CPF   string `json:"cpf" example:"123.456.789-00"`
}

// ToDomain converts UpdateCustomerRequest to domain.Customer
func (r UpdateCustomerRequest) ToDomain(id uint64) *domain.Customer {
	return &domain.Customer{
		ID:    id,
		Name:  r.Name,
		Email: r.Email,
		CPF:   r.CPF,
	}
}
