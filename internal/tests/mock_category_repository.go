package tests

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Insert(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(id uint64) (*domain.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetAll(name string, page, limit int) ([]domain.Category, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]domain.Category), args.Get(1).(int64), args.Error(2)
}
