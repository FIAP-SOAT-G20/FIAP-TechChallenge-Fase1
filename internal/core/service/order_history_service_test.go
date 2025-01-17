package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/mock/repository"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderHistoryService_Create(t *testing.T) {

	orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
	orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}

	t.Run("Should fail if order id is blank", func(t *testing.T) {
		orderHistoryRepositoryMock.ExpectedCalls = nil
		err := orderHistoryService.Create(0, nil, domain.UNDEFINDED)
		assert.EqualError(t, err, domain.ErrOrderIdMandatory.Error())
	})

	t.Run("Should create a new order history", func(t *testing.T) {
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.On("Insert", mock.Anything).Return(nil)
		err := orderHistoryService.Create(1, nil, domain.OPEN)
		assert.Nil(t, err)
	})
}

func TestOrderHistoryService_List(t *testing.T) {
	orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
	orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}

	t.Run("Should return an empty list if there is no ordert", func(t *testing.T) {

		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.On("GetAll", uint64(1), mock.Anything, 0, 10).Return(make([]domain.OrderHistory, 0), int64(0), nil)
		orders, size, err := orderHistoryService.List(uint64(1), nil, 0, 10)
		assert.Len(t, orders, 0)
		assert.Equal(t, int64(0), size)
		assert.Nil(t, err)
	})

	t.Run("Should return all orders history from given order", func(t *testing.T) {
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.On("GetAll", uint64(1), mock.Anything, 0, 10).Return([]domain.OrderHistory{{ID: 1, OrderID: 1, Status: domain.OPEN}, {ID: 2, OrderID: 1, Status: domain.READY}}, int64(2), nil)
		orders, size, err := orderHistoryService.List(uint64(1), nil, 0, 10)
		assert.Len(t, orders, 2)
		assert.Equal(t, int64(2), size)
		assert.Nil(t, err)
	})
}

func TestOrderHistoryService_Delete(t *testing.T) {
	orderHistoryRepositoryMock := &repository.OrderHistoryRepositoryMock{}
	orderHistoryService := OrderHistoryService{orderHistoryRepository: orderHistoryRepositoryMock}

	t.Run("Should fail if order doesn't exists", func(t *testing.T) {
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.On("GetByID", uint64(1)).Return((*domain.OrderHistory)(nil), domain.ErrNotFound)
		err := orderHistoryService.Delete(uint64(1))
		assert.NotNil(t, err)
		assert.EqualError(t, err, domain.ErrNotFound.Error())
	})

	t.Run("Should delete a order", func(t *testing.T) {
		orderHistoryRepositoryMock.ExpectedCalls = nil
		orderHistoryRepositoryMock.On("GetByID", uint64(1)).Return(&domain.OrderHistory{ID: 1, OrderID: 1}, nil)
		orderHistoryRepositoryMock.On("Delete", mock.AnythingOfType("uint64")).Return(nil)
		err := orderHistoryRepositoryMock.Delete(uint64(1))
		assert.Nil(t, err)
	})
}
