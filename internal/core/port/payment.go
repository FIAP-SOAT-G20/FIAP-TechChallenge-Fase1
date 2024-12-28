package port

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type IPaymentRepository interface {
	Insert(payment *domain.Payment) (*domain.Payment, error)
	GetPaymentByOrderIDAndStatus(status domain.PaymentStatus, orderID uint64) (*domain.Payment, error)
}

type IPaymentService interface {
	CreatePayment(orderID uint64) (*domain.Payment, error)
}

type IExternalPaymentService interface {
	CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error)
}
