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

func (ops *OrderProductService) Create(orderProduct *domain.OrderProduct) error {

	existingOrderProduct, err := ops.orderProductRepository.GetByID(orderProduct.OrderID, orderProduct.ProductID)
	if err != nil && err.Error() != "record not found" && err != domain.ErrNotFound {
		return err
	}
	if existingOrderProduct != nil {
		return domain.ErrConflict
	}

	order, err := ops.orderService.GetByID(orderProduct.OrderID)
	if err != nil {
		return domain.ErrOrderIdMandatory
	}
	if order.Status != domain.OPEN {
		return domain.ErrOrderIsNotOnStatusOpen
	}

	product, err := ops.productService.GetByID(orderProduct.ProductID)
	if err != nil {
		return domain.ErrProductIdMandatory
	}

	if orderProduct.Quantity <= 0 {
		return domain.ErrInvalidParam
	}

	orderProduct.Price = product.Price
	orderProduct.CreatedAt = time.Now()
	orderProduct.UpdatedAt = time.Now()

	err = ops.orderProductRepository.Insert(orderProduct)
	if err != nil {
		return err
	}
	totalBill, err := ops.orderProductRepository.GetTotalBillByOrderId(orderProduct.OrderID)
	if err != nil {
		return err
	}
	order.TotalBill = totalBill
	return ops.orderService.Update(order, nil)
}

func (ops *OrderProductService) GetByID(orderID, productID uint64) (*domain.OrderProduct, error) {
	orderProduct, err := ops.orderProductRepository.GetByID(orderID, productID)
	if err != nil {
		return nil, domain.ErrNotFound
	}
	return orderProduct, nil
}

func (ops *OrderProductService) List(orderID, productID uint64, page, limit int) ([]domain.OrderProduct, int64, error) {
	return ops.orderProductRepository.GetAll(orderID, productID, page, limit)
}

func (ops *OrderProductService) Update(orderProduct *domain.OrderProduct) error {
	existingOrderProduct, err := ops.orderProductRepository.GetByID(orderProduct.OrderID, orderProduct.ProductID)
	if err != nil {
		return domain.ErrNotFound
	}
	if orderProduct.Quantity <= 0 {
		return domain.ErrInvalidParam
	}
	order, err := ops.orderService.GetByID(orderProduct.OrderID)
	if err != nil {
		return domain.ErrNotFound
	}
	if order.Status != domain.OPEN {
		return domain.ErrOrderIsNotOnStatusOpen
	}
	orderProduct.Price = existingOrderProduct.Price
	orderProduct.UpdatedAt = time.Now()
	err = ops.orderProductRepository.Update(orderProduct)
	if err != nil {
		return err
	}
	totalBill, err := ops.orderProductRepository.GetTotalBillByOrderId(orderProduct.OrderID)
	if err != nil {
		return err
	}
	order.TotalBill = totalBill
	return ops.orderService.Update(order, nil)
}

func (ops *OrderProductService) Delete(orderID, productID uint64) error {
	_, err := ops.orderProductRepository.GetByID(orderID, productID)
	if err != nil {
		return domain.ErrNotFound
	}
	order, err := ops.orderService.GetByID(orderID)
	if err != nil {
		return domain.ErrNotFound
	}
	if order.Status != domain.OPEN {
		return domain.ErrOrderIsNotOnStatusOpen
	}
	err = ops.orderProductRepository.Delete(orderID, productID)
	if err != nil {
		return err
	}
	totalBill, err := ops.orderProductRepository.GetTotalBillByOrderId(orderID)
	if err != nil {
		return err
	}
	order.TotalBill = totalBill
	return ops.orderService.Update(order, nil)
}
