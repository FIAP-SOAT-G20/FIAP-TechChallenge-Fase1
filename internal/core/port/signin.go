package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type SigninService interface {
	GetByCPF(cpf string) (*domain.Customer, error)
}
