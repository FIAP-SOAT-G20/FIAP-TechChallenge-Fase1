package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (os *ProductServiceMock) Create(order *domain.Product) error {
	args := os.Called(order)
	return args.Error(0)
}

func (os *ProductServiceMock) GetByID(id uint64) (*domain.Product, error) {
	args := os.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (os *ProductServiceMock) List(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error) {
	args := os.Called(name, categoryID, page, limit)
	return args.Get(0).([]domain.Product), args.Get(1).(int64), args.Error(2)
}

func (os *ProductServiceMock) Update(order *domain.Product) error {
	args := os.Called(order)
	return args.Error(0)
}

func (os *ProductServiceMock) Delete(id uint64) error {
	args := os.Called(id)
	return args.Error(0)
}
