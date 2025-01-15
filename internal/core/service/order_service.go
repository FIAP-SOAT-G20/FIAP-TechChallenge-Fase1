package service

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderService struct {
	orderRepository     port.IOrderRepository
	orderHistoryService port.IOrderHistoryService
	customerService     port.ICustomerService
	staffService        port.IStaffService
}

func NewOrderService(orderRepository port.IOrderRepository, customerService port.ICustomerService, orderHistoryService port.IOrderHistoryService, staffService port.IStaffService) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		customerService:     customerService,
		orderHistoryService: orderHistoryService,
		staffService:        staffService,
	}
}

func (os *OrderService) Create(order *domain.Order) error {

	_, err := os.customerService.GetByID(order.CustomerID)
	if err != nil {
		return domain.ErrNotFound
	}

	order.Status = domain.OPEN
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	err = os.orderRepository.Insert(order)
	if err != nil {
		return err
	}

	return os.orderHistoryService.Create(order.ID, nil, order.Status)
}

func (os *OrderService) GetByID(id uint64) (*domain.Order, error) {
	return os.orderRepository.GetByID(id)
}

func (os *OrderService) List(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error) {
	return os.orderRepository.GetAll(customerID, status, page, limit)
}

func (os *OrderService) Update(order *domain.Order, staffID *uint64) error {
	existing, err := os.orderRepository.GetByID(order.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	if existing.CustomerID != order.CustomerID {
		return domain.ErrInvalidParam
	}

	if staffID != nil {
		_, err = os.staffService.GetByID(*staffID)
		if err != nil {
			return domain.ErrNotFound
		}
	}

	if existing.Status != order.Status && !domain.CanTransitionTo(existing.Status, order.Status) {
		return domain.ErrInvalidParam
	}

	order.UpdatedAt = time.Now()

	err = os.orderRepository.Update(order)
	if err != nil {
		return err
	}

	if existing.Status != order.Status {
		return os.orderHistoryService.Create(order.ID, staffID, order.Status)
	}
	return nil
}

func (os *OrderService) Delete(id uint64) error {
	_, err := os.orderRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return os.orderRepository.Delete(id)
}
