package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Insert(payment *domain.Payment) (*domain.Payment, error) {
	args := m.Called(payment)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Payment), args.Error(1)
}

func (m *MockPaymentRepository) GetPaymentByOrderIDAndStatus(status domain.PaymentStatus, orderID uint64) (*domain.Payment, error) {
	args := m.Called(status, orderID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Payment), args.Error(1)
}

func (m *MockPaymentRepository) UpdateStatus(status domain.PaymentStatus, externalPaymentID string) error {
	args := m.Called(status, externalPaymentID)
	return args.Error(0)
}

func (m *MockPaymentRepository) GetByExternalPaymentID(externalPaymentID string) (*domain.Payment, error) {
	args := m.Called(externalPaymentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Payment), args.Error(1)
}
