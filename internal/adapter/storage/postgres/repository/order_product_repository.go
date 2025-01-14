package repository

import (
	"errors"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"gorm.io/gorm"
)

type OrderProductRepository struct {
	db *gorm.DB
}

func NewOrderProductRepository(db *gorm.DB) *OrderProductRepository {
	return &OrderProductRepository{db}
}

func (r *OrderProductRepository) Insert(orderProduct *domain.OrderProduct) error {
	return r.db.Create(orderProduct).Error
}

func (r *OrderProductRepository) GetByID(id uint64) (*domain.OrderProduct, error) {
	var orderProduct domain.OrderProduct

	if err := r.db.
		Preload("Order").
		First(&orderProduct, id); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &orderProduct, nil
}

func (r *OrderProductRepository) GetAllByOrderID(orderID uint64) ([]domain.OrderProduct, error) {
	var orderProducts []domain.OrderProduct
	var tx = r.db.Model(&orderProducts).Preload("Order")
	where := map[string]interface{}{
		"order_id": orderID,
	}
	err := tx.Where(where).Find(&orderProducts).Error
	return orderProducts, err
}

func (r *OrderProductRepository) GetAll(orderID, productID *uint64, page, limit int) ([]domain.OrderProduct, int64, error) {
	var orderProducts []domain.OrderProduct
	var tx = r.db.Model(&orderProducts).Preload("Order")
	var count int64
	where := map[string]interface{}{}

	if orderID != nil {
		where["order_id"] = orderID
	}

	if productID != nil {
		where["product_id"] = productID
	}

	tx.Where(where).Count(&count)
	err := tx.Offset((page - 1) * limit).Limit(limit).Find(&orderProducts).Error
	return orderProducts, count, err
}

func (r *OrderProductRepository) GetTotalBillByOrderId(orderID uint64) (float32, error) {
	var total float32
	err := r.db.Model(&domain.OrderProduct{}).
		Select("sum(total_bill * quantity)").
		Where("order_id = ?", orderID).
		Scan(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderProductRepository) Update(orderProduct *domain.OrderProduct) error {
	return r.db.Save(orderProduct).Error
}

func (r *OrderProductRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.OrderProduct{}, id).Error
}
