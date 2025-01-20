package mocks

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockPaymentGatewayService struct {
	mock.Mock
}

func (m *MockPaymentGatewayService) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	args := m.Called(payment)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.CreatePaymentOUT), args.Error(1)
}
