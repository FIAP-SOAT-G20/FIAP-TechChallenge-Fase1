package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type StaffRepositoryMock struct {
	mock.Mock
}

func (repository *StaffRepositoryMock) Insert(staff *domain.Staff) error {
	args := repository.Called(staff)
	return args.Get(0).(error)

}

func (repository *StaffRepositoryMock) GetByID(id uint64) (*domain.Staff, error) {
	args := repository.Called(id)
	return args.Get(0).(*domain.Staff), args.Get(1).(error)
}

func (repository *StaffRepositoryMock) GetAll(name string, page, limit int) ([]domain.Staff, int64, error) {
	args := repository.Called(name, page, limit)
	return args.Get(0).([]domain.Staff), args.Get(1).(int64), args.Get(2).(error)
}

func (repository *StaffRepositoryMock) Update(staff *domain.Staff) error {
	args := repository.Called(staff)
	return args.Get(0).(error)
}

func (repository *StaffRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Get(0).(error)
}
