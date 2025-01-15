package service

import (
	"os"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PaymentService struct {
	paymentRepository      port.IPaymentRepository
	orderService           port.IOrderService
	externalPaymentService port.IExternalPaymentService
}

func NewPaymentService(
	paymentRepository port.IPaymentRepository,
	orderService port.IOrderService,
	externalPaymentService port.IExternalPaymentService,
) *PaymentService {
	return &PaymentService{
		paymentRepository:      paymentRepository,
		orderService:           orderService,
		externalPaymentService: externalPaymentService,
	}
}

func (ps *PaymentService) CreatePayment(orderID uint64) (*domain.Payment, error) {
	existentPedingPayment, err := ps.paymentRepository.GetPaymentByOrderIDAndStatus(domain.PROCESSING, orderID)
	if err != nil {
		return nil, err
	}

	if existentPedingPayment.ID != 0 {
		return existentPedingPayment, nil
	}

	order, err := ps.orderService.GetByID(orderID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	paymentPayload := ps.createPaymentPayload(order)

	extPayment, err := ps.externalPaymentService.CreatePaymentMock(paymentPayload)
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
	err = ps.orderService.Update(order, nil)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (ps *PaymentService) createPaymentPayload(order *domain.Order) *domain.CreatePaymentIN {
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
		NotificationUrl:   os.Getenv("MERCADO_PAGO_NOTIFICATION_URL"),
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
	err = ps.orderService.Update(order, nil)
	if err != nil {
		return nil, err
	}
	
	return paymentOUT, nil
}
