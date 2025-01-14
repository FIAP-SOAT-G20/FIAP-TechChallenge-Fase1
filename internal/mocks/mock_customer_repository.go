package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) Insert(customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetByID(id uint64) (*domain.Customer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetByCPF(cpf string) (*domain.Customer, error) {
	args := m.Called(cpf)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Customer), args.Error(1)
}

func (m *MockCustomerRepository) GetAll(name string, page, limit int) ([]domain.Customer, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]domain.Customer), args.Get(1).(int64), args.Error(2)
}

func (m *MockCustomerRepository) Update(customer *domain.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) Delete(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}
