package service

import (
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PaymentService struct {
	paymentRepository     port.IPaymentRepository
	orderService          port.IOrderService
	paymentGatewayService port.IPaymentGatewayService
}

func NewPaymentService(
	paymentRepository port.IPaymentRepository,
	orderService port.IOrderService,
	paymentGatewayService port.IPaymentGatewayService,
) *PaymentService {
	return &PaymentService{
		paymentRepository:     paymentRepository,
		orderService:          orderService,
		paymentGatewayService: paymentGatewayService,
	}
}

func (ps *PaymentService) CreatePayment(orderID uint64) (*domain.Payment, error) {
	existentPedingPayment, err := ps.paymentRepository.GetPaymentByOrderIDAndStatus(domain.PROCESSING, orderID)
	if err != nil && err != domain.ErrNotFound {
		return nil, err
	}

	if existentPedingPayment.ID != 0 {
		return existentPedingPayment, nil
	}

	order, err := ps.orderService.GetByID(orderID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	if len(order.OrderProducts) == 0 {
		return nil, domain.ErrOrderWithoutProducts
	}

	extPGPayload := createPaymentGatewayPayload(order)

	extPayment, err := ps.paymentGatewayService.CreatePayment(extPGPayload)
	if err != nil {
		return nil, err
	}

	iPayment := &domain.Payment{
		Status:            domain.PROCESSING,
		ExternalPaymentID: extPayment.InStoreOrderID,
		OrderID:           orderID,
		QrData:            extPayment.QrData,
	}

	payment, err := ps.paymentRepository.Insert(iPayment)
	if err != nil {
		return nil, err
	}

	order.Status = domain.PENDING
	err = ps.orderService.UpdateStatus(order, nil)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// remove ps
func createPaymentGatewayPayload(order *domain.Order) *domain.CreatePaymentIN {
	var items []domain.ItemsIN

	externalReference := strconv.FormatUint(order.ID, 10)

	for _, v := range order.OrderProducts {
		items = append(items, domain.ItemsIN{
			Title:       v.Product.Name,
			Description: v.Product.Description,
			UnitPrice:   float32(v.Product.Price),
			Category:    "marketplace",
			UnitMeasure: "unit",
			Quantity:    uint64(v.Quantity),
			TotalAmount: v.Price,
		})
	}

	return &domain.CreatePaymentIN{
		ExternalReference: externalReference,
		TotalAmount:       order.TotalBill,
		Items:             items,
		Title:             "FIAP Tech Challenge - Product Order",
		Description:       "Purchases made at the FIAP Tech Challenge store",
	}
}

func (ps *PaymentService) UpdatePayment(payment *domain.UpdatePaymentIN) (*domain.Payment, error) {
	if err := ps.paymentRepository.UpdateStatus(domain.CONFIRMED, payment.Resource); err != nil {
		return nil, err
	}

	paymentOUT, err := ps.paymentRepository.GetByExternalPaymentID(payment.Resource)
	if err != nil {
		return nil, err
	}

	order, err := ps.orderService.GetByID(paymentOUT.OrderID)
	if err != nil {
		return nil, err
	}

	order.Status = domain.RECEIVED
	err = ps.orderService.UpdateStatus(order, nil)
	if err != nil {
		return nil, err
	}

	return paymentOUT, nil
}
