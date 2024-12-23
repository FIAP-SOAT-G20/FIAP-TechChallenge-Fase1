package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) Insert(product *domain.Order) error {
	return r.db.Create(product).Error
}

func (r *OrderRepository) GetByID(id uint64) (*domain.Order, error) {
	var product domain.Order

	err := r.db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &product, err
}

func (r *OrderRepository) GetAll(name string, categoryID uint64, page, limit int) ([]domain.Order, int64, error) {
	var products []domain.Order
	var count int64

	query := r.db.Model(&domain.Order{})

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

func (r *OrderRepository) Update(product *domain.Order) error {
	return r.db.Save(product).Error
}

func (r *OrderRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
