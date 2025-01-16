package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/mocks"
)

func TestOrderService_Create(t *testing.T) {
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockOrderHistoryRepo := new(mocks.MockOrderHistoryRepository)
	mockOrderProductRepo := new(mocks.MockOrderProductRepository)
	mockCustomerRepo := new(mocks.MockCustomerRepository)
	orderService := NewOrderService(mockOrderRepo, mockOrderHistoryRepo, mockOrderProductRepo, mockCustomerRepo)

	scenarios := []struct {
		name          string
		order         *domain.Order
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given valid order When Create is called Then should succeed",
			order: &domain.Order{
				CustomerID: 1,
				TotalBill:  100.0,
				OrderProducts: []domain.OrderProduct{
					{
						ProductID: 1,
						Quantity:  2,
						Price:     50.0,
					},
				},
			},
			setupMocks: func() {
				mockCustomerRepo.On("GetByID", uint64(1)).Return(&domain.Customer{ID: 1}, nil)
				mockOrderRepo.On("Insert", mock.AnythingOfType("*domain.Order")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non-existent customer When Create is called Then should return error",
			order: &domain.Order{
				CustomerID: 1,
				TotalBill:  100.0,
			},
			setupMocks: func() {
				mockCustomerRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
		{
			name: "Given invalid total bill When Create is called Then should return error",
			order: &domain.Order{
				CustomerID: 1,
				TotalBill:  0,
			},
			setupMocks: func() {
				mockCustomerRepo.On("GetByID", uint64(1)).Return(&domain.Customer{ID: 1}, nil)
			},
			expectedError: domain.ErrInvalidParam,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockOrderRepo.ExpectedCalls = nil
				mockOrderHistoryRepo.ExpectedCalls = nil
				mockOrderProductRepo.ExpectedCalls = nil
				mockCustomerRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := orderService.Create(tt.order)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestOrderService_GetByID(t *testing.T) {
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockOrderHistoryRepo := new(mocks.MockOrderHistoryRepository)
	mockOrderProductRepo := new(mocks.MockOrderProductRepository)
	mockCustomerRepo := new(mocks.MockCustomerRepository)
	orderService := NewOrderService(mockOrderRepo, mockOrderHistoryRepo, mockOrderProductRepo, mockCustomerRepo)

	scenarios := []struct {
		name          string
		id            uint64
		setupMocks    func()
		expectedOrder *domain.Order
		expectedError error
	}{
		{
			name: "Given existing order ID When GetByID is called Then should return order",
			id:   1,
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1}, nil)
			},
			expectedOrder: &domain.Order{ID: 1},
			expectedError: nil,
		},
		{
			name: "Given non-existent order ID When GetByID is called Then should return error",
			id:   1,
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedOrder: nil,
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockOrderRepo.ExpectedCalls = nil
				mockOrderHistoryRepo.ExpectedCalls = nil
				mockOrderProductRepo.ExpectedCalls = nil
				mockCustomerRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			order, err := orderService.GetByID(tt.id)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOrder, order)
		})
	}
}

func TestOrderService_List(t *testing.T) {
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockOrderHistoryRepo := new(mocks.MockOrderHistoryRepository)
	mockOrderProductRepo := new(mocks.MockOrderProductRepository)
	mockCustomerRepo := new(mocks.MockCustomerRepository)
	orderService := NewOrderService(mockOrderRepo, mockOrderHistoryRepo, mockOrderProductRepo, mockCustomerRepo)

	scenarios := []struct {
		name          string
		clientID      uint64
		page          int
		limit         int
		setupMocks    func()
		expectedCount int64
		expectedError error
	}{
		{
			name:     "Given valid parameters When List is called Then should return orders",
			clientID: 1,
			page:     1,
			limit:    10,
			setupMocks: func() {
				orders := []domain.Order{{ID: 1, CustomerID: 1}}
				mockOrderRepo.On("GetAll", uint64(1), 1, 10).Return(orders, int64(1), nil)
			},
			expectedCount: 1,
			expectedError: nil,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockOrderRepo.ExpectedCalls = nil
				mockOrderHistoryRepo.ExpectedCalls = nil
				mockOrderProductRepo.ExpectedCalls = nil
				mockCustomerRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			_, count, err := orderService.List(tt.clientID, tt.page, tt.limit)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedCount, count)
		})
	}
}

func TestOrderService_Update(t *testing.T) {
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockOrderHistoryRepo := new(mocks.MockOrderHistoryRepository)
	mockOrderProductRepo := new(mocks.MockOrderProductRepository)
	mockCustomerRepo := new(mocks.MockCustomerRepository)
	orderService := NewOrderService(mockOrderRepo, mockOrderHistoryRepo, mockOrderProductRepo, mockCustomerRepo)

	currentTime := time.Now()

	scenarios := []struct {
		name          string
		order         *domain.Order
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given existing order When Update is called Then should succeed",
			order: &domain.Order{
				ID:         1,
				CustomerID: 1,
				CreatedAt:  currentTime,
			},
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1}, nil)
				mockOrderRepo.On("Update", mock.AnythingOfType("*domain.Order")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non-existent order When Update is called Then should return error",
			order: &domain.Order{
				ID:         1,
				CustomerID: 1,
			},
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockOrderRepo.ExpectedCalls = nil
				mockOrderHistoryRepo.ExpectedCalls = nil
				mockOrderProductRepo.ExpectedCalls = nil
				mockCustomerRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := orderService.Update(tt.order)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestOrderService_Delete(t *testing.T) {
	mockOrderRepo := new(mocks.MockOrderRepository)
	mockOrderHistoryRepo := new(mocks.MockOrderHistoryRepository)
	mockOrderProductRepo := new(mocks.MockOrderProductRepository)
	mockCustomerRepo := new(mocks.MockCustomerRepository)
	orderService := NewOrderService(mockOrderRepo, mockOrderHistoryRepo, mockOrderProductRepo, mockCustomerRepo)

	scenarios := []struct {
		name          string
		id            uint64
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given existing order When Delete is called Then should succeed",
			id:   1,
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(&domain.Order{ID: 1}, nil)
				mockOrderRepo.On("Delete", uint64(1)).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non-existent order When Delete is called Then should return error",
			id:   1,
			setupMocks: func() {
				mockOrderRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockOrderRepo.ExpectedCalls = nil
				mockOrderHistoryRepo.ExpectedCalls = nil
				mockOrderProductRepo.ExpectedCalls = nil
				mockCustomerRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := orderService.Delete(tt.id)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
