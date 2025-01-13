package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/mock/repository"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderService_Create(t *testing.T) {
	t.Run("Should fail if customer does not exist", func(t *testing.T) {
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("Create", mock.AnythingOfType("Order")).Return(nil)

		customerRepositoryMock := &repository.CustomerRepositoryMock{}
		customerRepositoryMock.On("GetByID", uint64(1)).Return((*domain.Customer)(nil), domain.ErrNotFound)

		orderService := OrderService{orderRepository: orderRepositoryMock, customerRepository: customerRepositoryMock}
		err := orderService.Create(&domain.Order{CustomerID: 1})
		assert.NotNil(t, err)
	})

	t.Run("Should create a new order", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1}
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("Insert", mock.Anything).Return(nil)

		customerRepositoryMock := &repository.CustomerRepositoryMock{}
		customerRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Customer{ID: 1}, nil)

		orderService := OrderService{orderRepository: orderRepositoryMock, customerRepository: customerRepositoryMock}
		err := orderService.Create(&order)
		assert.Nil(t, err)
	})
}

func TestOrderService_List(t *testing.T) {
	t.Run("Should return an empty list if there is no order from client", func(t *testing.T) {

		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetAll", uint64(1), 0, 10).Return(make([]domain.Order, 0), int64(0), nil)

		orderService := OrderService{orderRepository: orderRepositoryMock}
		orders, size, err := orderService.List(1, 0, 10)
		assert.Len(t, orders, 0)
		assert.Equal(t, int64(0), size)
		assert.Nil(t, err)
	})

	t.Run("Should return all orders from client", func(t *testing.T) {
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetAll", uint64(1), 0, 10).Return([]domain.Order{{ID: 1, CustomerID: 1}, {ID: 2, CustomerID: 1}}, int64(2), nil)

		orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
		orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}

		orderService := OrderService{orderRepository: orderRepositoryMock, orderHistoryService: &orderHistoryService}
		orders, size, err := orderService.List(uint64(1), 0, 10)
		assert.Len(t, orders, 2)
		assert.Equal(t, int64(2), size)
		assert.Nil(t, err)
	})
}

func TestOrderService_Update(t *testing.T) {

	t.Run("Should fail if customer has changed", func(t *testing.T) {
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1, CustomerID: 1}, nil)

		orderService := OrderService{orderRepository: orderRepositoryMock}
		err := orderService.Update(&domain.Order{ID: 1, CustomerID: 2})
		assert.NotNil(t, err)
	})

	t.Run("Should update a order", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1}

		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&order, nil)
		orderRepositoryMock.On("Update", &order).Return(nil)

		orderService := OrderService{orderRepository: orderRepositoryMock}
		err := orderService.Update(&order)
		assert.Nil(t, err)
	})
}

func TestOrderService_Delete(t *testing.T) {
	t.Run("Should fail if order doesn't exists", func(t *testing.T) {
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetByID", uint64(1)).Return((*domain.Order)(nil), domain.ErrNotFound)

		orderService := OrderService{orderRepository: orderRepositoryMock}
		err := orderService.Delete(uint64(1))
		assert.NotNil(t, err)
	})

	t.Run("Should delete a order", func(t *testing.T) {
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1, CustomerID: 1}, nil)
		orderRepositoryMock.On("Delete", mock.AnythingOfType("uint64")).Return(nil)

		orderService := OrderService{orderRepository: orderRepositoryMock}
		err := orderService.Delete(uint64(1))
		assert.Nil(t, err)
	})
}
