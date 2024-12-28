package service

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderService struct {
	orderRepository        port.IOrderRepository
	orderHistoryRepository port.IOrderHistoryRepository
	orderProductRepository port.IOrderProductRepository
	customerRepository     port.ICustomerRepository
}

func NewOrderService(orderRepository port.IOrderRepository, customerRepository port.ICustomerRepository) *OrderService {
	return &OrderService{
		orderRepository:    orderRepository,
		customerRepository: customerRepository,
	}
}

func (ps *OrderService) Create(order *domain.Order) error {

	_, err := ps.customerRepository.GetByID(order.CustomerID)
	if err != nil {
		return domain.ErrNotFound
	}

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	err = ps.orderRepository.Insert(order)
	if err != nil {
		return err
	}

	// TODO throw event order created
	return nil
}

func (ps *OrderService) GetByID(id uint64) (*domain.Order, error) {
	return ps.orderRepository.GetByID(id)
}

func (ps *OrderService) List(clientId uint64, page, limit int) ([]domain.Order, int64, error) {
	return ps.orderRepository.GetAll(clientId, page, limit)
}

func (ps *OrderService) Update(order *domain.Order) error {
	existing, err := ps.orderRepository.GetByID(order.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	if existing.CustomerID != order.CustomerID {
		return domain.ErrInvalidParam
	}

	order.UpdatedAt = time.Now()

	return ps.orderRepository.Update(order)
}

func (ps *OrderService) Delete(id uint64) error {
	_, err := ps.orderRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return ps.orderRepository.Delete(id)
}
