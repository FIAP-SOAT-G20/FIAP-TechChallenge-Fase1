package service

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderService struct {
	orderRepository        port.IOrderRepository
	orderHistoryService    port.IOrderHistoryService
	orderProductRepository port.IOrderProductRepository
	customerRepository     port.ICustomerRepository
}

func NewOrderService(orderRepository port.IOrderRepository, customerRepository port.ICustomerRepository, orderHistoryService port.IOrderHistoryService) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		customerRepository:  customerRepository,
		orderHistoryService: orderHistoryService,
	}
}

func (os *OrderService) Create(order *domain.Order) error {

	_, err := os.customerRepository.GetByID(order.CustomerID)
	if err != nil {
		return domain.ErrNotFound
	}

	order.Status = domain.RECEIVED
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	err = os.orderRepository.Insert(order)
	if err != nil {
		return err
	}

	return os.orderHistoryService.Create(order.ID, nil, domain.RECEIVED)
}

func (os *OrderService) GetByID(id uint64) (*domain.Order, error) {
	return os.orderRepository.GetByID(id)
}

func (os *OrderService) List(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error) {
	return os.orderRepository.GetAll(customerID, status, page, limit)
}

func (os *OrderService) Update(order *domain.Order) error {
	existing, err := os.orderRepository.GetByID(order.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	if existing.CustomerID != order.CustomerID {
		return domain.ErrInvalidParam
	}

	order.UpdatedAt = time.Now()

	return os.orderRepository.Update(order)
}

func (os *OrderService) Delete(id uint64) error {
	_, err := os.orderRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return os.orderRepository.Delete(id)
}
