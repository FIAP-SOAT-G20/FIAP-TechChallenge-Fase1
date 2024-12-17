package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"time"
)

type OrderService struct {
	orderRepository        port.OrderRepository
	orderHistoryRepository port.OrderHistoryRepository
	orderProductRepository port.OrderProductRepository
	customerRepository     port.CustomerRepository
}

func NewOrderService(orderRepository port.OrderRepository, orderHistoryRepo port.OrderHistoryRepository, orderProductRepo port.OrderProductRepository, customerRepository port.CustomerRepository) *OrderService {
	return &OrderService{
		orderRepository:        orderRepository,
		orderHistoryRepository: orderHistoryRepo,
		orderProductRepository: orderProductRepo,
		customerRepository:     customerRepository,
	}
}

func (ps *OrderService) Create(order *domain.Order) error {

	_, err := ps.customerRepository.GetByID(order.CustomerID)
	if err != nil {
		return domain.ErrNotFound
	}

	if order.TotalBill <= 0 {
		return domain.ErrInvalidParam
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
	_, err := ps.orderRepository.GetByID(order.ID)
	if err != nil {
		return domain.ErrNotFound
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
