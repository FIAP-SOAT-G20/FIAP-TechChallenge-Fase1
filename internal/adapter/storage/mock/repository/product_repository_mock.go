package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (repository *ProductRepositoryMock) Insert(product *domain.Product) error {
	args := repository.Called(product)
	return args.Error(0)

}

func (repository *ProductRepositoryMock) GetByID(id uint64) (*domain.Product, error) {
	args := repository.Called(id)
	if args.Get(0) == nil || args.Get(0) == uint64(0) {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (repository *ProductRepositoryMock) GetAll(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error) {
	args := repository.Called(name, categoryID, page, limit)
	return args.Get(0).([]domain.Product), args.Get(1).(int64), args.Error(2)
}

func (repository *ProductRepositoryMock) Update(staff *domain.Product) error {
	args := repository.Called(staff)
	return args.Error(0)
}

func (repository *ProductRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Error(0)
}
