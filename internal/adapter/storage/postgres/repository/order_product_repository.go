package repository

import (
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

func (r *OrderProductRepository) GetByID(orderID, productID uint64) (*domain.OrderProduct, error) {
	var item domain.OrderProduct
	err := r.db.Where(&domain.OrderProduct{OrderID: orderID, ProductID: productID}).Preload("Order").Preload("Product").First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *OrderProductRepository) GetAllByOrderID(orderID uint64) ([]domain.OrderProduct, error) {
	var orderProducts []domain.OrderProduct
	var tx = r.db.Model(&orderProducts).Preload("Order").Preload("Product")
	where := map[string]interface{}{
		"order_id": orderID,
	}
	err := tx.Where(where).Find(&orderProducts).Error
	return orderProducts, err
}

func (r *OrderProductRepository) GetAll(orderID, productID uint64, page, limit int) ([]domain.OrderProduct, int64, error) {
	var orderProducts []domain.OrderProduct
	var tx = r.db.Model(&orderProducts).Preload("Order").Preload("Product")
	var count int64
	where := map[string]interface{}{}

	if orderID > 0 {
		where["order_id"] = orderID
	}

	if productID > 0 {
		where["product_id"] = productID
	}

	tx.Where(where).Count(&count)
	err := tx.Offset((page - 1) * limit).Limit(limit).Find(&orderProducts).Error
	return orderProducts, count, err
}

func (r *OrderProductRepository) GetTotalBillByOrderId(orderID uint64) (float32, error) {
	var total float32
	err := r.db.Model(&domain.OrderProduct{}).
		Select("sum(price * quantity)").
		Where("order_id = ?", orderID).
		Scan(&total).Error

	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderProductRepository) Update(orderProduct *domain.OrderProduct) error {
	return r.db.
		Where("order_id = ? and product_id = ?", orderProduct.OrderID, orderProduct.ProductID).
		Updates(&domain.OrderProduct{
			Quantity:  orderProduct.Quantity,
			UpdatedAt: orderProduct.UpdatedAt,
		}).Error
}

func (r *OrderProductRepository) Delete(orderID, productID uint64) error {
	return r.db.Model(&domain.OrderProduct{}).
		Where("order_id = ? and product_id = ? ", orderID, productID).
		Delete(&domain.OrderProduct{
			OrderID:   orderID,
			ProductID: productID,
		}).Error
}
