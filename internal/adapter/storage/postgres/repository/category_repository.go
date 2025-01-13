package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (r *CategoryRepository) Insert(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetByID(id uint64) (*domain.Category, error) {
	var category domain.Category

	err := r.db.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &category, err
}

func (r *CategoryRepository) GetAll(name string, page, limit int) ([]domain.Category, int64, error) {
	var categories []domain.Category
	var count int64

	query := r.db.Model(&domain.Category{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Count(&count)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&categories).Error

	return categories, count, err
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Category{}, id).Error
}
