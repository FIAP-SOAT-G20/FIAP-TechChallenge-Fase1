package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type SignInService struct {
	customerRepository  port.CustomerRepository
}

func NewSignInService(customerRepository  port.CustomerRepository) *SignInService {
	return &SignInService{
		customerRepository: customerRepository,
	}
}

func (ps *SignInService) GetByCPF(cpf string)  (*domain.Customer, error) {
	return ps.customerRepository.GetByCPF(cpf)
}

