package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderHistoryRepositoryMock struct {
	mock.Mock
}

func (repository *OrderHistoryRepositoryMock) Insert(orderHistory *domain.OrderHistory) error {
	args := repository.Called(orderHistory)
	return args.Error(0)

}

func (repository *OrderHistoryRepositoryMock) GetByID(id uint64) (*domain.OrderHistory, error) {
	args := repository.Called(id)
	return args.Get(0).(*domain.OrderHistory), args.Error(1)
}

func (repository *OrderHistoryRepositoryMock) GetAll(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error) {
	args := repository.Called(orderID, status, page, limit)
	return args.Get(0).([]domain.OrderHistory), args.Get(1).(int64), args.Error(2)
}

func (repository *OrderHistoryRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Error(0)
}
