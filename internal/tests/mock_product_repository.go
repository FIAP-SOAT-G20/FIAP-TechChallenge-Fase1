package tests

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Insert(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(id uint64) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepository) GetAll(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error) {
	args := m.Called(name, categoryID, page, limit)
	return args.Get(0).([]domain.Product), args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) Update(product *domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}
