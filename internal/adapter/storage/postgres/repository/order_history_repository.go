package repository

import (
	"errors"
	"fmt"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"gorm.io/gorm"
)

type OrderHistoryRepository struct {
	db *gorm.DB
}

func NewOrderHistoryRepository(db *gorm.DB) *OrderHistoryRepository {
	return &OrderHistoryRepository{db}
}

func (r *OrderHistoryRepository) Insert(orderHistory *domain.OrderHistory) error {
	return r.db.Create(orderHistory).Error
}

func (r *OrderHistoryRepository) GetByID(id uint64) (*domain.OrderHistory, error) {
	var orderHistory domain.OrderHistory

	if err := r.db.
		Preload("Order").
		First(&orderHistory, id); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &orderHistory, nil
}

func (r *OrderHistoryRepository) GetAll(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error) {
	// Add the orderHistories get all
	var orderHistories []domain.OrderHistory
	var tx = r.db.Model(&orderHistories).Preload("Order")
	var count int64
	where := map[string]interface{}{}
	fmt.Printf("Parameters: \n")
	fmt.Printf("OrderHistoryID: %d\n", orderID)
	fmt.Printf("Status: %v\n", status.ToString())
	fmt.Printf("Page: %d\n", page)
	fmt.Printf("Limit: %d\n", limit)
	if &status != nil && *status != domain.UNDEFINDED {
		tx = r.db.Preload("OrderHistory")
		where["status"] = &status
	}

	if orderID > 0 {
		where["order_id"] = orderID
	}

	tx.Where(where).Count(&count)

	err := tx.Offset((page - 1) * limit).Limit(limit).Find(&orderHistories).Error

	return orderHistories, count, err
}

func (r *OrderHistoryRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.OrderHistory{}, id).Error
}
