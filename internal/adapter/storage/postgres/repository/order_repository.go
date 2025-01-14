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

func (r *OrderRepository) Insert(order *domain.Order) error {
	return r.db.Create(order).Error
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

func (r *OrderRepository) GetAll(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error) {
	var orders []domain.Order
	var tx = r.db
	var count int64

	var filter = &domain.Order{}
	if &status != nil && *status != domain.UNDEFINDED {
		filter.Status = *status
	}

	if customerID > 0 {
		filter.CustomerID = customerID
	}

	err := tx.Model(filter).Offset((page-1)*limit).Limit(limit).Find(&orders, filter).Count(&count).Error
	return orders, count, err
}

func (r *OrderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
