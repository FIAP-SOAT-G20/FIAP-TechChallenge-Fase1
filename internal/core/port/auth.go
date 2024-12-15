package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type TokenService interface {
	CreateToken(customer *domain.Customer) (string, error)
	VerifyToken(token string) (*domain.TokenPayload, error)
}

type AuthService interface {
	Login(email, password string) (string, error)
}
