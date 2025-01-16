package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockExternalPaymentRepository struct {
	mock.Mock
}

func (m *MockExternalPaymentRepository) Insert(payment *domain.Payment) (*domain.Payment, error) {
	args := m.Called(payment)
	return args.Get(0).(*domain.Payment), args.Error(1)
}
