package service

import (
	"errors"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/mock/repository"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderProductService_Create(t *testing.T) {

	orderServiceMock := new(OrderServiceMock)
	productServiceMock := new(ProductServiceMock)
	orderProductRepositoryMock := new(repository.OrderProductRepositoryMock)
	orderProductService := NewOrderProductService(orderProductRepositoryMock, orderServiceMock, productServiceMock)

	resetMocks := func() {
		orderServiceMock.ExpectedCalls = nil
		productServiceMock.ExpectedCalls = nil
		orderProductRepositoryMock.ExpectedCalls = nil
	}

	t.Run("Should fail if order not found", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 0, ProductID: 0, Quantity: 1}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(nil, domain.ErrOrderIdMandatory)
		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrOrderIdMandatory.Error())
	})

	t.Run("Should fail if product not found", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 0, Quantity: 1}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1}, nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return((*domain.Product)(nil), errors.New("product not found"))
		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrProductIdMandatory.Error())
	})

	t.Run("Should fail if quantity is lower or equal than zero", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 1, Quantity: 0}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1}, nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Product{ID: 1}, nil)

		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrInvalidParam.Error())
	})

	t.Run("Should create a new order product", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 0, Quantity: 1}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN}, nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Product{ID: 1, Price: 1}, nil)
		orderProductRepositoryMock.On("GetTotalBillByOrderId", mock.AnythingOfType("uint64")).Return(float32(1), nil)
		orderServiceMock.On("Update", mock.Anything, mock.Anything).Return(nil)
		orderProductRepositoryMock.On("Insert", mock.Anything).Return(nil)

		err := orderProductService.Create(&orderProduct)
		assert.Nil(t, err)
		assert.Equal(t, float32(1), orderProduct.Price)
	})
}

func TestOrderProductService_Update(t *testing.T) {

	orderServiceMock := new(OrderServiceMock)
	productServiceMock := new(ProductServiceMock)
	orderProductRepositoryMock := new(repository.OrderProductRepositoryMock)
	orderProductService := NewOrderProductService(orderProductRepositoryMock, orderServiceMock, productServiceMock)

	resetMocks := func() {
		orderServiceMock.ExpectedCalls = nil
		productServiceMock.ExpectedCalls = nil
		orderProductRepositoryMock.ExpectedCalls = nil
	}

	t.Run("Should fail if order not found", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 0, ProductID: 0, Quantity: 1}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(nil, domain.ErrOrderIdMandatory)
		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrOrderIdMandatory.Error())
	})

	t.Run("Should fail if product not found", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 0, Quantity: 1}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1}, nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return((*domain.Product)(nil), errors.New("product not found"))
		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrProductIdMandatory.Error())
	})

	t.Run("Should fail if quantity is lower or equal than zero", func(t *testing.T) {
		resetMocks()
		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 1, Quantity: 0}
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1}, nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Product{ID: 1}, nil)

		err := orderProductService.Create(&orderProduct)
		assert.EqualError(t, err, domain.ErrInvalidParam.Error())
	})

	t.Run("Should update a order product", func(t *testing.T) {
		resetMocks()
		orderServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Order{ID: 1, CustomerID: 1, Status: domain.OPEN}, nil)
		orderServiceMock.On("Update", mock.Anything, mock.Anything).Return(nil)
		productServiceMock.On("GetByID", mock.AnythingOfType("uint64")).Return(&domain.Product{ID: 1, Price: 1}, nil)
		orderProductRepositoryMock.On("GetByID", uint64(1), uint64(1)).Return(&domain.OrderProduct{OrderID: 1, ProductID: 1, Quantity: 1, Price: 1}, nil)
		orderProductRepositoryMock.On("GetTotalBillByOrderId", mock.AnythingOfType("uint64")).Return(float32(2), nil)
		orderProductRepositoryMock.On("Update", mock.Anything).Return(nil)

		orderProduct := domain.OrderProduct{OrderID: 1, ProductID: 1, Quantity: 2, Price: 1}
		err := orderProductService.Update(&orderProduct)
		assert.Nil(t, err)
		assert.Equal(t, float32(1), orderProduct.Price)
	})
}

func TestOrderProductService_List(t *testing.T) {
	orderServiceMock := new(OrderServiceMock)
	productRepositoryMock := new(repository.ProductRepositoryMock)
	categoryRepositoryMock := new(repository.CategoryRepositoryMock)
	productService := NewProductService(productRepositoryMock, categoryRepositoryMock)
	orderProductRepositoryMock := &repository.OrderProductRepositoryMock{}
	orderProductService := NewOrderProductService(orderProductRepositoryMock, orderServiceMock, productService)

	resetMocks := func() {
		productRepositoryMock.ExpectedCalls = nil
		orderProductRepositoryMock.ExpectedCalls = nil
	}

	t.Run("Should return an empty list if there is no order product", func(t *testing.T) {

		resetMocks()
		orderProductRepositoryMock.On("GetAll", uint64(0), uint64(0), 0, 10).Return(make([]domain.OrderProduct, 0), int64(0), nil)
		orders, size, err := orderProductService.List(uint64(0), uint64(0), 0, 10)
		assert.Len(t, orders, 0)
		assert.Equal(t, int64(0), size)
		assert.Nil(t, err)
	})

	t.Run("Should return all orders products from given order", func(t *testing.T) {
		resetMocks()
		orderProductRepositoryMock.On("GetAll", uint64(1), uint64(0), 0, 10).Return([]domain.OrderProduct{{OrderID: 1, ProductID: 1, Price: 1, Quantity: 1}, {OrderID: 1, ProductID: 2, Price: 2, Quantity: 1}}, int64(2), nil)
		orders, size, err := orderProductService.List(uint64(1), uint64(0), 0, 10)
		assert.Len(t, orders, 2)
		assert.Equal(t, int64(2), size)
		assert.Nil(t, err)
	})
}

func TestOrderProductService_Delete(t *testing.T) {
	orderServiceMock := new(OrderServiceMock)
	productServiceMock := new(ProductServiceMock)
	orderProductRepositoryMock := &repository.OrderProductRepositoryMock{}
	orderProductService := NewOrderProductService(orderProductRepositoryMock, orderServiceMock, productServiceMock)

	resetMocks := func() {
		orderServiceMock.ExpectedCalls = nil
		productServiceMock.ExpectedCalls = nil
		orderProductRepositoryMock.ExpectedCalls = nil
	}

	t.Run("Should fail if order doesn't exists", func(t *testing.T) {
		resetMocks()
		orderProductRepositoryMock.On("GetByID", uint64(1), uint64(1)).Return(&domain.OrderProduct{OrderID: 1, ProductID: 1}, nil)
		orderServiceMock.On("GetByID", uint64(1)).Return((*domain.Order)(nil), domain.ErrNotFound)
		err := orderProductService.Delete(uint64(1), uint64(1))
		assert.NotNil(t, err)
		assert.EqualError(t, err, domain.ErrNotFound.Error())
	})

	t.Run("Should fail if order product doesn't exists", func(t *testing.T) {
		resetMocks()
		orderProductRepositoryMock.On("GetByID", uint64(1), uint64(0)).Return((*domain.OrderProduct)(nil), domain.ErrNotFound)
		orderServiceMock.On("GetByID", uint64(1)).Return(domain.Order{ID: 1}, domain.ErrNotFound)
		err := orderProductService.Delete(uint64(1), uint64(0))
		assert.NotNil(t, err)
		assert.EqualError(t, err, domain.ErrNotFound.Error())
	})

	t.Run("Should delete a order product", func(t *testing.T) {
		resetMocks()
		orderProductRepositoryMock.On("GetByID", uint64(1), uint64(1)).Return(&domain.OrderProduct{OrderID: 1, ProductID: 1, Price: 1, Quantity: 1}, nil)
		orderProductRepositoryMock.On("GetTotalBillByOrderId", mock.AnythingOfType("uint64")).Return(float32(0), nil)
		orderProductRepositoryMock.On("Delete", uint64(1), uint64(1)).Return(nil)
		orderServiceMock.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1, Status: domain.OPEN}, nil)
		orderServiceMock.On("Update", mock.Anything, mock.Anything).Return(nil)
		err := orderProductService.Delete(uint64(1), uint64(1))
		assert.Nil(t, err)
	})
}
