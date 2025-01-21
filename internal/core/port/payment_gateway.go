package port

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type IPaymentGatewayRepository interface {
	CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error)
}
