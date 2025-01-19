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

		orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
		orderHistoryRepositoryMock.On("Insert", mock.AnythingOfType("OrderHistory")).Return(nil)

		orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}
		customerService := CustomerService{customerRepository: customerRepositoryMock}
		orderService := OrderService{orderRepository: orderRepositoryMock, customerService: &customerService, orderHistoryService: &orderHistoryService}
		err := orderService.Create(&domain.Order{CustomerID: 1})
		assert.NotNil(t, err)
	})

	t.Run("Should create a new order", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1}
		orderRepositoryMock := &repository.OrderRepositoryMock{}
		orderRepositoryMock.On("Insert", mock.Anything).Return(nil)

		orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
		orderHistoryRepositoryMock.On("Insert", mock.Anything).Return(nil)
		orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}

		customerRepositoryMock := &repository.CustomerRepositoryMock{}
		customerRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Customer{ID: 1}, nil)
		customerService := CustomerService{customerRepository: customerRepositoryMock}
		orderService := OrderService{orderRepository: orderRepositoryMock, customerService: &customerService, orderHistoryService: &orderHistoryService}
		err := orderService.Create(&order)
		assert.Nil(t, err)
	})
}

func TestOrderService_List(t *testing.T) {
	orderRepositoryMock := &repository.OrderRepositoryMock{}
	orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
	orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}
	orderService := OrderService{orderRepository: orderRepositoryMock, orderHistoryService: &orderHistoryService}

	t.Run("Should return an empty list if there is no order from client", func(t *testing.T) {
		orderRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetAll", uint64(1), mock.Anything, 0, 10).Return(make([]domain.Order, 0), int64(0), nil)
		orders, size, err := orderService.List(uint64(1), nil, 0, 10)
		assert.Len(t, orders, 0)
		assert.Equal(t, int64(0), size)
		assert.Nil(t, err)
	})

	t.Run("Should return all orders from client", func(t *testing.T) {
		orderRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetAll", uint64(1), mock.Anything, 0, 10).Return([]domain.Order{{ID: 1, CustomerID: 1}, {ID: 2, CustomerID: 1}}, int64(2), nil)
		orders, size, err := orderService.List(uint64(1), nil, 0, 10)
		assert.Len(t, orders, 2)
		assert.Equal(t, int64(2), size)
		assert.Nil(t, err)
	})
}

func TestOrderService_UpdateStatus(t *testing.T) {
	orderRepositoryMock := &repository.OrderRepositoryMock{}
	orderService := OrderService{orderRepository: orderRepositoryMock}
	t.Run("Should fail if customer has changed", func(t *testing.T) {
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN}, nil)
		err := orderService.UpdateStatus(&domain.Order{ID: 1, CustomerID: 2, Status: domain.OPEN}, nil)
		assert.NotNil(t, err)
	})

	t.Run("Should update a order status", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN}
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&order, nil)
		orderRepositoryMock.On("UpdateStatus", &order).Return(nil)
		err := orderService.UpdateStatus(&order, nil)
		assert.Nil(t, err)
	})

	t.Run("Should fail if status transition not allowed", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN}
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&order, nil)
		orderToUpdate := domain.Order{ID: 1, CustomerID: 1, Status: domain.COMPLETED}
		err := orderService.UpdateStatus(&orderToUpdate, nil)
		assert.EqualError(t, err, domain.ErrOrderInvalidStatusTransition.Error())
	})

	t.Run("Should fail if can transition but staff if not informed", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1, Status: domain.PREPARING}
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&order, nil)
		orderToUpdate := domain.Order{ID: 1, CustomerID: 1, Status: domain.READY}
		err := orderService.UpdateStatus(&orderToUpdate, nil)
		assert.EqualError(t, err, domain.ErrOrderMandatoryStaffId.Error())
	})

	t.Run("Should fail if status changed to PENDING without Products", func(t *testing.T) {
		order := domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN, OrderProducts: make([]domain.OrderProduct, 0)}
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&order, nil)
		orderToUpdate := domain.Order{ID: 1, CustomerID: 1, Status: domain.PENDING, OrderProducts: make([]domain.OrderProduct, 0)}
		err := orderService.UpdateStatus(&orderToUpdate, nil)
		assert.EqualError(t, err, domain.ErrOrderWithoutProducts.Error())
	})

}

func TestOrderService_Delete(t *testing.T) {
	orderRepositoryMock := &repository.OrderRepositoryMock{}
	orderService := OrderService{orderRepository: orderRepositoryMock}

	t.Run("Should fail if order doesn't exists", func(t *testing.T) {
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return((*domain.Order)(nil), domain.ErrNotFound)
		err := orderService.Delete(uint64(1))
		assert.NotNil(t, err)
	})

	t.Run("Should delete a order", func(t *testing.T) {
		orderRepositoryMock.ExpectedCalls = nil
		orderRepositoryMock.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1, CustomerID: 1}, nil)
		orderRepositoryMock.On("Delete", mock.AnythingOfType("uint64")).Return(nil)
		err := orderService.Delete(uint64(1))
		assert.Nil(t, err)
	})
}
