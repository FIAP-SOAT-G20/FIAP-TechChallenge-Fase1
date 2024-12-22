package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/tests"
)

func TestProductService_Create(t *testing.T) {
	mockProductRepo := new(tests.MockProductRepository)
	mockCategoryRepo := new(tests.MockCategoryRepository)
	productService := NewProductService(mockProductRepo, mockCategoryRepo)

	scenarios := []struct {
		name          string
		product       *domain.Product
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given valid product When Create is called Then should succeed",
			product: &domain.Product{
				Name:       "Test Product",
				CategoryID: 1,
			},
			setupMocks: func() {
				category := &domain.Category{ID: 1, Name: "Test Category"}
				mockCategoryRepo.On("GetByID", uint64(1)).Return(category, nil)
				mockProductRepo.On("Insert", mock.AnythingOfType("*domain.Product")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given valid product When repository fails Then should return error",
			product: &domain.Product{
				Name:       "Test Product",
				CategoryID: 1,
			},
			setupMocks: func() {
				category := &domain.Category{ID: 1, Name: "Test Category"}
				mockCategoryRepo.On("GetByID", uint64(1)).Return(category, nil)
				mockProductRepo.On("Insert", mock.AnythingOfType("*domain.Product")).Return(domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockProductRepo.ExpectedCalls = nil
				mockCategoryRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := productService.Create(tt.product)
			assert.Equal(t, tt.expectedError, err)
			mockProductRepo.AssertExpectations(t)
			mockCategoryRepo.AssertExpectations(t)
		})
	}
}

func TestProductService_GetByID(t *testing.T) {
	mockProductRepo := new(tests.MockProductRepository)
	mockCategoryRepo := new(tests.MockCategoryRepository)
	productService := NewProductService(mockProductRepo, mockCategoryRepo)

	scenarios := []struct {
		name            string
		id              uint64
		setupMocks      func()
		expectedProduct *domain.Product
		expectedError   error
	}{
		{
			name: "Given existing product ID When GetByID is called Then should return product",
			id:   1,
			setupMocks: func() {
				product := &domain.Product{ID: 1, Name: "Test Product"}
				mockProductRepo.On("GetByID", uint64(1)).Return(product, nil)
			},
			expectedProduct: &domain.Product{ID: 1, Name: "Test Product"},
			expectedError:   nil,
		},
		{
			name: "Given non existing product ID When GetByID is called Then should return not found",
			id:   1,
			setupMocks: func() {
				mockProductRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedProduct: nil,
			expectedError:   domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockProductRepo.ExpectedCalls = nil
				mockCategoryRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			product, err := productService.GetByID(tt.id)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedProduct, product)
			mockProductRepo.AssertExpectations(t)
		})
	}
}

func TestProductService_List(t *testing.T) {
	mockProductRepo := new(tests.MockProductRepository)
	mockCategoryRepo := new(tests.MockCategoryRepository)
	productService := NewProductService(mockProductRepo, mockCategoryRepo)

	scenarios := []struct {
		name             string
		searchName       string
		categoryID       uint64
		page             int
		limit            int
		setupMocks       func()
		expectedProducts []domain.Product
		expectedTotal    int64
		expectedError    error
	}{
		{
			name:       "Given valid search criteria When List is called Then should return filtered products",
			searchName: "Test",
			categoryID: 1,
			page:       1,
			limit:      10,
			setupMocks: func() {
				products := []domain.Product{{ID: 1, Name: "Test Product"}}
				mockProductRepo.On("GetAll", "Test", uint64(1), 1, 10).Return(products, int64(1), nil)
			},
			expectedProducts: []domain.Product{{ID: 1, Name: "Test Product"}},
			expectedTotal:    1,
			expectedError:    nil,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockProductRepo.ExpectedCalls = nil
				mockCategoryRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			products, total, err := productService.List(tt.searchName, tt.categoryID, tt.page, tt.limit)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedProducts, products)
			assert.Equal(t, tt.expectedTotal, total)
			mockProductRepo.AssertExpectations(t)
		})
	}
}

func TestProductService_Update(t *testing.T) {
	mockProductRepo := new(tests.MockProductRepository)
	mockCategoryRepo := new(tests.MockCategoryRepository)
	productService := NewProductService(mockProductRepo, mockCategoryRepo)

	scenarios := []struct {
		name          string
		product       *domain.Product
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given existing product When Update is called Then should succeed",
			product: &domain.Product{
				ID:   1,
				Name: "Updated Product",
			},
			setupMocks: func() {
				mockProductRepo.On("GetByID", uint64(1)).Return(&domain.Product{ID: 1}, nil)
				mockProductRepo.On("Update", mock.AnythingOfType("*domain.Product")).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non existing product When Update is called Then should return not found",
			product: &domain.Product{
				ID:   1,
				Name: "Updated Product",
			},
			setupMocks: func() {
				mockProductRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockProductRepo.ExpectedCalls = nil
				mockCategoryRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := productService.Update(tt.product)
			assert.Equal(t, tt.expectedError, err)
			mockProductRepo.AssertExpectations(t)
		})
	}
}

func TestProductService_Delete(t *testing.T) {
	mockProductRepo := new(tests.MockProductRepository)
	mockCategoryRepo := new(tests.MockCategoryRepository)
	productService := NewProductService(mockProductRepo, mockCategoryRepo)

	scenarios := []struct {
		name          string
		id            uint64
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given existing product ID When Delete is called Then should succeed",
			id:   1,
			setupMocks: func() {
				mockProductRepo.On("GetByID", uint64(1)).Return(&domain.Product{ID: 1}, nil)
				mockProductRepo.On("Delete", uint64(1)).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Given non existing product ID When Delete is called Then should return not found",
			id:   1,
			setupMocks: func() {
				mockProductRepo.On("GetByID", uint64(1)).Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockProductRepo.ExpectedCalls = nil
				mockCategoryRepo.ExpectedCalls = nil
			})

			tt.setupMocks()
			err := productService.Delete(tt.id)
			assert.Equal(t, tt.expectedError, err)
			mockProductRepo.AssertExpectations(t)
		})
	}
}
