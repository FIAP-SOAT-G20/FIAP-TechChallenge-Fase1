package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Insert(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetByID(id uint64) (*domain.Product, error) {
	var product domain.Product

	err := r.db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &product, err
}

func (r *ProductRepository) GetAll(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error) {
	var products []domain.Product
	var count int64

	query := r.db.Model(&domain.Product{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	query.Count(&count)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&products).Error

	return products, count, err
}

func (r *ProductRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Product{}, id).Error
}
