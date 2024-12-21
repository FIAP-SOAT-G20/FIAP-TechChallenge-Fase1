package response

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type SignInResponse struct {
	CustomerResponse
}

func NewSignInResponse(customer *domain.Customer) SignInResponse {
	if customer == nil {
		return SignInResponse{}
	}

	return SignInResponse{
		CustomerResponse: CustomerResponse{
			ID:    customer.ID,
			Name:  customer.Name,
			Email: customer.Email,
			CPF:   customer.CPF,
		},
	}
}
