package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type ProductService struct {
	productRepository  port.IProductRepository
	categoryRepository port.ICategoryRepository
}

func NewProductService(productRepository port.IProductRepository, categoryRepository port.ICategoryRepository) *ProductService {
	return &ProductService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}

func (ps *ProductService) Create(product *domain.Product) error {
	category, err := ps.categoryRepository.GetByID(product.CategoryID)
	if err != nil {
		return domain.ErrNotFound
	}

	product.Category = *category
	product.Active = true

	return ps.productRepository.Insert(product)
}

func (ps *ProductService) GetByID(id uint64) (*domain.Product, error) {
	return ps.productRepository.GetByID(id)
}

func (ps *ProductService) List(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error) {
	return ps.productRepository.GetAll(name, categoryID, page, limit)
}

func (ps *ProductService) Update(product *domain.Product) error {
	_, err := ps.productRepository.GetByID(product.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	return ps.productRepository.Update(product)
}

func (ps *ProductService) Delete(id uint64) error {
	_, err := ps.productRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return ps.productRepository.Delete(id)
}
