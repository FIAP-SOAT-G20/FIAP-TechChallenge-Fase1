package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (repository *OrderRepositoryMock) Insert(order *domain.Order) error {
	args := repository.Called(order)
	return args.Error(0)

}

func (repository *OrderRepositoryMock) GetByID(id uint64) (*domain.Order, error) {
	args := repository.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (repository *OrderRepositoryMock) GetAll(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error) {
	args := repository.Called(customerID, status, page, limit)
	return args.Get(0).([]domain.Order), args.Get(1).(int64), args.Error(2)
}

func (repository *OrderRepositoryMock) Update(order *domain.Order) error {
	args := repository.Called(order)
	return args.Error(0)
}

func (repository *OrderRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Error(0)
}
