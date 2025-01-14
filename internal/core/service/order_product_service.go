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

func (os *OrderProductService) Create(orderProduct domain.OrderProduct) error {

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

	err = os.orderProductRepository.Insert(&orderProduct)
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

func (ps *OrderProductService) List(orderID, productID *uint64, page, limit int) ([]domain.OrderProduct, int64, error) {

	if orderID != nil && *orderID <= 0 {
		return nil, 0, domain.ErrInvalidParam
	}
	if productID != nil && *productID <= 0 {
		return nil, 0, domain.ErrInvalidParam
	}

	return ps.orderProductRepository.GetAll(orderID, productID, page, limit)
}

func (os *OrderProductService) Update(orderProduct *domain.OrderProduct) error {
	existing, err := os.orderProductRepository.GetByID(orderProduct.ID)
	if err != nil {
		return domain.ErrNotFound
	}
	if existing.OrderID != orderProduct.OrderID {
		return domain.ErrInvalidParam
	}
	if existing.ProductID != orderProduct.ProductID {
		return domain.ErrInvalidParam
	}
	if orderProduct.Quantity <= 0 {
		return domain.ErrInvalidParam
	}
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

func (os *OrderProductService) Delete(id uint64) error {
	_, err := os.orderProductRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return os.orderProductRepository.Delete(id)
}
