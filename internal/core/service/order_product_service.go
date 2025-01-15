package service

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderProductService struct {
	orderProductRepository port.IOrderProductRepository
	orderService           port.IOrderService
	productService         port.IProductService
}

func NewOrderProductService(orderProductRepository port.IOrderProductRepository, orderService port.IOrderService, productService port.IProductService) *OrderProductService {
	return &OrderProductService{
		orderProductRepository: orderProductRepository,
		orderService:           orderService,
		productService:         productService,
	}
}

func (os *OrderProductService) Create(orderProduct *domain.OrderProduct) error {

	order, err := os.orderService.GetByID(orderProduct.OrderID)
	if err != nil {
		return domain.ErrNotFound
	}

	product, err := os.productService.GetByID(orderProduct.ProductID)
	if err != nil {
		return domain.ErrNotFound
	}

	if orderProduct.Quantity <= 0 {
		return domain.ErrInvalidParam
	}

	orderProduct.Price = product.Price
	orderProduct.CreatedAt = time.Now()
	orderProduct.UpdatedAt = time.Now()

	err = os.orderProductRepository.Insert(orderProduct)
	if err != nil {
		return err
	}
	totalBill, err := os.orderProductRepository.GetTotalBillByOrderId(orderProduct.OrderID)
	if err != nil {
		return err
	}
	order.TotalBill = totalBill
	return os.orderService.Update(order, nil)
}

func (ps *OrderProductService) GetByID(id uint64) (*domain.OrderProduct, error) {
	return ps.orderProductRepository.GetByID(id)
}

func (ps *OrderProductService) List(orderID, productID uint64, page, limit int) ([]domain.OrderProduct, int64, error) {
	return ps.orderProductRepository.GetAll(orderID, productID, page, limit)
}

func (os *OrderProductService) Update(orderProduct *domain.OrderProduct) error {
	orderProducts, _, err := os.orderProductRepository.GetAll(orderProduct.OrderID, orderProduct.ProductID, 0, 1)
	if err != nil || len(orderProducts) == 0 {
		return domain.ErrNotFound
	}
	if orderProduct.Quantity <= 0 {
		return domain.ErrInvalidParam
	}
	orderProduct.Price = orderProducts[0].Price
	orderProduct.UpdatedAt = time.Now()
	err = os.orderProductRepository.Update(orderProduct)
	if err != nil {
		return err
	}
	totalBill, err := os.orderProductRepository.GetTotalBillByOrderId(orderProduct.OrderID)
	if err != nil {
		return err
	}
	order, err := os.orderService.GetByID(orderProduct.OrderID)
	if err != nil {
		return domain.ErrNotFound
	}
	order.TotalBill = totalBill
	return os.orderService.Update(order, nil)
}

func (os *OrderProductService) Delete(orderProduct *domain.OrderProduct) error {
	orderProducts, _, err := os.orderProductRepository.GetAll(orderProduct.OrderID, orderProduct.ProductID, 0, 1)
	if err != nil || len(orderProducts) == 0 {
		return domain.ErrNotFound
	}

	return os.orderProductRepository.Delete(&orderProducts[0])
}
