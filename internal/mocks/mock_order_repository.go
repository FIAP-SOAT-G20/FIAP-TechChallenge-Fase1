package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Insert(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetByID(id uint64) (*domain.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (m *MockOrderRepository) GetAll(name uint64, page, limit int) ([]domain.Order, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]domain.Order), args.Get(1).(int64), args.Error(2)
}

func (m *MockOrderRepository) Update(order *domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) Delete(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}
