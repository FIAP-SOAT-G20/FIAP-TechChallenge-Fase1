package service

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderHistoryService struct {
	orderHistoryRepository port.IOrderHistoryRepository
}

func NewOrderHistoryService(orderHistoryRepository port.IOrderHistoryRepository) *OrderHistoryService {
	return &OrderHistoryService{
		orderHistoryRepository: orderHistoryRepository,
	}
}

func (ohs *OrderHistoryService) Create(orderID uint64, staffID *uint64, status domain.OrderStatus) error {

	if orderID == 0 {
		return domain.ErrOrderIdMandatory
	}

	orderHistory := domain.OrderHistory{
		OrderID:   orderID,
		StaffID:   staffID,
		Status:    status,
		CreatedAt: time.Now(),
	}

	return ohs.orderHistoryRepository.Insert(&orderHistory)
}

func (ohs *OrderHistoryService) GetByID(id uint64) (*domain.OrderHistory, error) {
	return ohs.orderHistoryRepository.GetByID(id)
}

func (ohs *OrderHistoryService) List(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error) {
	return ohs.orderHistoryRepository.GetAll(orderID, status, page, limit)
}

func (ohs *OrderHistoryService) Delete(id uint64) error {
	_, err := ohs.orderHistoryRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return ohs.orderHistoryRepository.Delete(id)
}
