package repository

import (
	"errors"
	"fmt"
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
	// Add the order get all
	var orders []domain.Order
	var tx = r.db
	var count int64

	var filter = &domain.Order{}
	fmt.Printf("Parameters--------------------: \n")
	fmt.Printf("CustomerID: %d\n", customerID)
	fmt.Printf("Status: %v\n", status.ToString())
	fmt.Printf("Page: %d\n", page)
	fmt.Printf("Limit: %d\n", limit)
	fmt.Printf("Antes do if do status\n")

	if &status != nil && *status != domain.UNDEFINDED {
		filter.Status = *status
	}

	if customerID > 0 {
		filter.CustomerID = customerID
	}

	/*fmt.Printf("Chegou no count\n")
	tx.Count(&count)*/

	fmt.Printf("Passou do where e do count")
	err := tx.Model(filter).Offset((page-1)*limit).Limit(limit).Find(&orders, filter).Count(&count).Error

	fmt.Printf("Passou do find. Count: %d, Error: %v", count, err)
	return orders, count, err
}

func (r *OrderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Order{}, id).Error
}
