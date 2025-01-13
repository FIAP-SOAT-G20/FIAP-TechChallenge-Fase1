package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type CategoryService struct {
	categoryRepository port.ICategoryRepository
}

func NewCategoryService(categoryRepository port.ICategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (cs *CategoryService) Create(category *domain.Category) error {
	return cs.categoryRepository.Insert(category)
}

func (cs *CategoryService) GetByID(id uint64) (*domain.Category, error) {
	return cs.categoryRepository.GetByID(id)
}

func (cs *CategoryService) List(name string, page, limit int) ([]domain.Category, int64, error) {
	return cs.categoryRepository.GetAll(name, page, limit)
}

func (cs *CategoryService) Update(category *domain.Category) error {
	_, err := cs.categoryRepository.GetByID(category.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	return cs.categoryRepository.Update(category)
}

func (cs *CategoryService) Delete(id uint64) error {
	_, err := cs.categoryRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return cs.categoryRepository.Delete(id)
}
