package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderProductRepositoryMock struct {
	mock.Mock
}

func (repository *OrderProductRepositoryMock) Insert(orderProduct *domain.OrderProduct) error {
	args := repository.Called(orderProduct)
	return args.Error(0)

}

func (repository *OrderProductRepositoryMock) GetByID(orderID, productID uint64) (*domain.OrderProduct, error) {
	args := repository.Called(orderID, productID)
	if args.Get(0) == 0 || args.Get(1) == 0 {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.OrderProduct), args.Error(1)
}

func (repository *OrderProductRepositoryMock) GetAllByOrderID(orderID uint64) ([]domain.OrderProduct, error) {
	args := repository.Called(orderID)
	if args.Get(0) == 0 {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.OrderProduct), args.Error(1)
}

func (repository *OrderProductRepositoryMock) GetAll(orderID, productID uint64, page, limit int) ([]domain.OrderProduct, int64, error) {
	args := repository.Called(orderID, productID, page, limit)
	return args.Get(0).([]domain.OrderProduct), args.Get(1).(int64), args.Error(2)
}

func (repository *OrderProductRepositoryMock) GetTotalBillByOrderId(orderID uint64) (float32, error) {
	args := repository.Called(orderID)
	return args.Get(0).(float32), args.Error(1)
}

func (repository *OrderProductRepositoryMock) Update(order *domain.OrderProduct) error {
	args := repository.Called(order)
	return args.Error(0)
}

func (repository *OrderProductRepositoryMock) Delete(orderID, productID uint64) error {
	args := repository.Called(orderID, productID)
	return args.Error(0)
}
