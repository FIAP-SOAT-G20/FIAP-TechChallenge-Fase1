package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (repository *CategoryRepositoryMock) Insert(staff *domain.Category) error {
	args := repository.Called(staff)
	return args.Get(0).(error)

}

func (repository *CategoryRepositoryMock) GetByID(id uint64) (*domain.Category, error) {
	args := repository.Called(id)
	return args.Get(0).(*domain.Category), args.Get(1).(error)
}

func (repository *CategoryRepositoryMock) GetAll(name string, page, limit int) ([]domain.Category, int64, error) {
	args := repository.Called(name, page, limit)
	return args.Get(0).([]domain.Category), args.Get(1).(int64), args.Get(2).(error)
}

func (repository *CategoryRepositoryMock) Update(staff *domain.Category) error {
	args := repository.Called(staff)
	return args.Get(0).(error)
}

func (repository *CategoryRepositoryMock) Delete(id uint64) error {
	args := repository.Called(id)
	return args.Get(0).(error)
}
