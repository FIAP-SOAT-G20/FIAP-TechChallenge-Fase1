package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type ISignInService interface {
	GetByCPF(cpf string) (*domain.Customer, error)
}
