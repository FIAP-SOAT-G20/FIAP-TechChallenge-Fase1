package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (repository *CustomerRepositoryMock) Insert(order *domain.Customer) error {
	args := repository.Called(order)
	return args.Get(0).(error)

}

func (repository *CustomerRepositoryMock) GetByID(id uint64) (*domain.Customer, error) {
	args := repository.Called(id)
	customerArg := args.Get(0)
	if customerArg == nil {
		return nil, args.Error(1)
	}
	errorArg := args.Get(1)
	if errorArg == nil {
		return args.Get(0).(*domain.Customer), nil
	}
	return args.Get(0).(*domain.Customer), args.Get(1).(error)
}

func (repository *CustomerRepositoryMock) GetByCPF(cpf string) (*domain.Customer, error) {
	args := repository.Called(cpf)
	return args.Get(0).(*domain.Customer), args.Get(1).(error)
}

func (repository *CustomerRepositoryMock) GetAll(name string, page, limit int) ([]domain.Customer, int64, error) {
	args := repository.Called(name, page, limit)
	return args.Get(0).([]domain.Customer), args.Get(1).(int64), args.Get(2).(error)
}

func (repository *CustomerRepositoryMock) Update(order *domain.Customer) error {
	args := repository.Called(order)
	return args.Get(0).(error)
}

func (repository *CustomerRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Get(0).(error)
}
