package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockOrderHistoryRepository struct {
	mock.Mock
}

func (m *MockOrderHistoryRepository) Insert(orderHistory *domain.OrderHistory) error {
	args := m.Called(orderHistory)
	return args.Error(0)
}
