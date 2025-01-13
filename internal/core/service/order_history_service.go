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

func (os *OrderHistoryService) Create(orderID uint64, staffID *uint64, status domain.OrderStatus) error {

	orderHistory := domain.OrderHistory{
		OrderID:   orderID,
		StaffID:   staffID,
		Status:    status,
		CreatedAt: time.Now(),
	}

	return os.orderHistoryRepository.Insert(&orderHistory)
}

func (ps *OrderHistoryService) GetByID(id uint64) (*domain.OrderHistory, error) {
	return ps.orderHistoryRepository.GetByID(id)
}

func (ps *OrderHistoryService) List(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error) {

	if orderID <= 0 {
		return nil, 0, domain.ErrInvalidParam
	}

	return ps.orderHistoryRepository.GetAll(orderID, status, page, limit)
}

func (ps *OrderHistoryService) Delete(id uint64) error {
	_, err := ps.orderHistoryRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return ps.orderHistoryRepository.Delete(id)
}
