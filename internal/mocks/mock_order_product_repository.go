package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockOrderProductRepository struct {
	mock.Mock
}

func (m *MockOrderProductRepository) Insert(orderProduct *domain.OrderProduct) error {
	args := m.Called(orderProduct)
	return args.Error(0)
}

func (m *MockOrderProductRepository) InsertMany(orderProducts []domain.OrderProduct) error {
	args := m.Called(orderProducts)
	return args.Error(0)
}
