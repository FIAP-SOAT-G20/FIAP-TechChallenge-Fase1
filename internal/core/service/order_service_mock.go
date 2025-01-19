package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type OrderServiceMock struct {
	mock.Mock
}

func (os *OrderServiceMock) Create(order *domain.Order) error {
	args := os.Called(order)
	return args.Error(0)

}

func (os *OrderServiceMock) GetByID(id uint64) (*domain.Order, error) {
	args := os.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Order), args.Error(1)
}

func (os *OrderServiceMock) List(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error) {
	args := os.Called(customerID, status, page, limit)
	return args.Get(0).([]domain.Order), args.Get(1).(int64), args.Error(2)
}

func (os *OrderServiceMock) UpdateStatus(order *domain.Order, staffID *uint64) error {
	args := os.Called(order, staffID)
	return args.Error(0)
}

func (os *OrderServiceMock) Delete(id uint64) error {
	args := os.Called(id)
	return args.Error(0)
}
