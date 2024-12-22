package repository

import (
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) GetByID(id uint64) (*domain.Order, error) {
	var order domain.Order

	if err := r.db.
		Preload("Customer").
		Preload("OrderProducts.Product").
		First(&order, id); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &order, nil
}

func (r *OrderRepository) Insert(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetAll(clientID uint64, page, limit int) ([]domain.Order, int64, error) {
	// Add the order get all
	return nil, 0, nil
}

func (r *OrderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
