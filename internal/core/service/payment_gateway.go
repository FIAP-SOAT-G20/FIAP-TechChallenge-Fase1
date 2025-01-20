package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PaymentGatewayService struct {
	paymentGatewayRepository port.IPaymentGatewayRepository
}

func NewPaymentGatewayService(paymentGatewayRepository port.IPaymentGatewayRepository) *PaymentGatewayService {
	return &PaymentGatewayService{
		paymentGatewayRepository: paymentGatewayRepository,
	}
}

func (pgs *PaymentGatewayService) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	return pgs.paymentGatewayRepository.CreatePayment(payment)
}
