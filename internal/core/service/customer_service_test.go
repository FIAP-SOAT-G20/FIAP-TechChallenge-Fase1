package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/tests"
)

func TestCustomerService_Create(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		customer      *domain.Customer
		setupMocks    func()
		expectedError error
	}{
		{
			name:     "Given valid customer When Create is called Then should succeed",
			customer: tests.MockCustomer(),
			setupMocks: func() {
				mockCustomerRepository.On("Insert", mock.AnythingOfType("*domain.Customer")).Return(nil)
			},
			expectedError: nil,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := customerService.Create(tt.customer)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCustomerService_GetByID(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		id            uint64
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given valid ID When GetByID is called Then should succeed",
			id:   1,
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", uint64(1)).Return(tests.MockCustomer(), nil)
			},
			expectedError: nil,
		},
		{
			name: "Given invalid ID When GetByID is called Then should return error",
			id:   1,
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			_, err := customerService.GetByID(tt.id)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCustomerService_GetByCPF(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		cpf           string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given valid CPF When GetByCPF is called Then should succeed",
			cpf:  "12345678900",
			setupMocks: func() {
				mockCustomerRepository.On("GetByCPF", "12345678900").Return(tests.MockCustomer(), nil)
			},
			expectedError: nil,
		},
		{
			name: "Given invalid CPF When GetByCPF is called Then should return error",
			cpf:  "12345678900",
			setupMocks: func() {
				mockCustomerRepository.On("GetByCPF", "12345678900").Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			_, err := customerService.GetByCPF(tt.cpf)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCustomerService_List(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		searchName    string
		page, limit   int
		setupMocks    func()
		expectedError error
	}{
		{
			name:       "Given valid params When List is called Then should succeed",
			searchName: "John",
			page:       1,
			limit:      10,
			setupMocks: func() {
				mockCustomerRepository.On("GetAll", "John", 1, 10).Return([]domain.Customer{*tests.MockCustomer()}, int64(1), nil)
			},
			expectedError: nil,
		},
		{
			name:       "Given no results When List is called Then should return empty list",
			searchName: "NonExistent",
			page:       1,
			limit:      10,
			setupMocks: func() {
				mockCustomerRepository.On("GetAll", "NonExistent", 1, 10).Return([]domain.Customer{}, int64(0), nil)
			},
			expectedError: nil,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			_, _, err := customerService.List(tt.searchName, tt.page, tt.limit)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCustomerService_Update(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		customer      *domain.Customer
		setupMocks    func()
		expectedError error
	}{
		{
			name:     "Given existing customer When Update is called Then should succeed",
			customer: tests.MockCustomer(),
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", tests.MockCustomer().ID).Return(tests.MockCustomer(), nil)
				mockCustomerRepository.On("Update", mock.AnythingOfType("*domain.Customer")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:     "Given non-existent customer When Update is called Then should return error",
			customer: tests.MockCustomer(),
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", tests.MockCustomer().ID).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := customerService.Update(tt.customer)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCustomerService_Delete(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	customerService := NewCustomerService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		id            uint64
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given existing customer ID When Delete is called Then should succeed",
			id:   1,
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", uint64(1)).Return(tests.MockCustomer(), nil)
				mockCustomerRepository.On("Delete", uint64(1)).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non-existent customer ID When Delete is called Then should return error",
			id:   1,
			setupMocks: func() {
				mockCustomerRepository.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := customerService.Delete(tt.id)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
