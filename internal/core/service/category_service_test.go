package service

import (
	"errors"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Insert(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) GetByID(id uint64) (*domain.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetAll(name string, page, limit int) ([]domain.Category, int64, error) {
	args := m.Called(name, page, limit)
	return args.Get(0).([]domain.Category), args.Get(1).(int64), args.Error(2)
}

func (m *MockCategoryRepository) Update(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCategoryService_Create(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	t.Run("sucesso ao criar categoria", func(t *testing.T) {
		category := &domain.Category{Name: "Test Category"}
		mockRepo.On("Insert", category).Return(nil)

		err := service.Create(category)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao criar categoria", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		category := &domain.Category{Name: "Test Category"}
		expectedErr := errors.New("erro ao inserir")
		mockRepo.On("Insert", mock.AnythingOfType("*domain.Category")).Return(expectedErr)

		err := service.Create(category)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_GetByID(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	t.Run("sucesso ao buscar categoria", func(t *testing.T) {
		expectedCategory := &domain.Category{ID: 1, Name: "Test Category"}
		mockRepo.On("GetByID", uint64(1)).Return(expectedCategory, nil)

		category, err := service.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategory, category)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao buscar categoria", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		expectedErr := errors.New("categoria não encontrada")
		mockRepo.On("GetByID", uint64(1)).Return(nil, expectedErr)

		category, err := service.GetByID(1)

		assert.Error(t, err)
		assert.Nil(t, category)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_List(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	t.Run("sucesso ao listar categorias", func(t *testing.T) {
		expectedCategories := []domain.Category{{ID: 1, Name: "Category 1"}}
		expectedTotal := int64(1)
		mockRepo.On("GetAll", "test", 1, 10).Return(expectedCategories, expectedTotal, nil)

		categories, total, err := service.List("test", 1, 10)

		assert.NoError(t, err)
		assert.Equal(t, expectedCategories, categories)
		assert.Equal(t, expectedTotal, total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao listar categorias", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		expectedErr := errors.New("erro ao listar")
		var emptyCategories []domain.Category
		mockRepo.On("GetAll", "test", 1, 10).Return(emptyCategories, int64(0), expectedErr)

		categories, total, err := service.List("test", 1, 10)

		assert.Error(t, err)
		assert.Empty(t, categories)
		assert.Equal(t, int64(0), total)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_Update(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	t.Run("sucesso ao atualizar categoria", func(t *testing.T) {
		category := &domain.Category{ID: 1, Name: "Updated Category"}
		mockRepo.On("GetByID", uint64(1)).Return(category, nil)
		mockRepo.On("Update", category).Return(nil)

		err := service.Update(category)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro quando categoria não existe", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		category := &domain.Category{ID: 1, Name: "Updated Category"}
		mockRepo.On("GetByID", uint64(1)).Return(nil, errors.New("não encontrado"))

		err := service.Update(category)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrNotFound, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao atualizar categoria", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		category := &domain.Category{ID: 1, Name: "Updated Category"}
		expectedErr := errors.New("erro ao atualizar")
		mockRepo.On("GetByID", uint64(1)).Return(category, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Category")).Return(expectedErr)

		err := service.Update(category)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCategoryService_Delete(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	service := NewCategoryService(mockRepo)

	t.Run("sucesso ao deletar categoria", func(t *testing.T) {
		category := &domain.Category{ID: 1, Name: "Category"}
		mockRepo.On("GetByID", uint64(1)).Return(category, nil)
		mockRepo.On("Delete", uint64(1)).Return(nil)

		err := service.Delete(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro quando categoria não existe", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		mockRepo.On("GetByID", uint64(1)).Return(nil, errors.New("não encontrado"))

		err := service.Delete(1)

		assert.Error(t, err)
		assert.Equal(t, domain.ErrNotFound, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("erro ao deletar categoria", func(t *testing.T) {
		mockRepo := new(MockCategoryRepository)
		service := NewCategoryService(mockRepo)
		
		category := &domain.Category{ID: 1, Name: "Category"}
		expectedErr := errors.New("erro ao deletar")
		mockRepo.On("GetByID", uint64(1)).Return(category, nil)
		mockRepo.On("Delete", uint64(1)).Return(expectedErr)

		err := service.Delete(1)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
} 