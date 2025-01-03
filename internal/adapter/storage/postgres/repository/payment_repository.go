package repository

import (
	"errors"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) Insert(payment *domain.Payment) (*domain.Payment, error) {
	if err := r.db.Create(payment).Error; err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) GetPaymentByOrderIDAndStatus(status domain.PaymentStatus, orderID uint64) (*domain.Payment, error) {
	var payment domain.Payment

	if err := r.db.Where("order_id = ? AND status = ?", orderID, status).First(&payment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &payment, nil
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) UpdateStatus(status domain.PaymentStatus, externalPaymentID string) error {
	if err := r.db.Model(&domain.Payment{}).Where("external_payment_id = ?", externalPaymentID).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
func (r *PaymentRepository) GetByExternalPaymentID(externalPaymentID string) (*domain.Payment, error) {
	var payment domain.Payment

	if err := r.db.Where("external_payment_id = ?", externalPaymentID).First(&payment); errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &payment, nil
}
