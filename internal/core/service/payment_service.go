package service

import (
	"os"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PaymentService struct {
	paymentRepository      port.IPaymentRepository
	orderRepository        port.IOrderRepository
	externalPaymentService port.IExternalPaymentService
}

func NewPaymentService(
	paymentRepository port.IPaymentRepository,
	orderRepository port.IOrderRepository,
	externalPaymentService port.IExternalPaymentService,
) *PaymentService {
	return &PaymentService{
		paymentRepository:      paymentRepository,
		orderRepository:        orderRepository,
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

	order, err := ps.orderRepository.GetByID(orderID)
	if err != nil {
		return nil, domain.ErrNotFound
	}

	paymentPayload := ps.createPaymentPayload(order)

	mpPayment, err := ps.externalPaymentService.CreatePayment(paymentPayload)
	if err != nil {
		return nil, err
	}

	iPayment := &domain.Payment{
		Status:            domain.PROCESSING,
		ExternalPaymentID: mpPayment.InStoreOrderID,
		OrderID:           orderID,
		QrData:            mpPayment.QrData,
	}

	payment, err := ps.paymentRepository.Insert(iPayment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (ps *PaymentService) createPaymentPayload(order *domain.Order) *request.CreatePaymentRequest {
	var items []request.Items

	externalReference := strconv.FormatUint(order.ID, 10)

	for _, v := range order.OrderProducts {
		items = append(items, request.Items{
			Title:       v.Product.Name,
			Description: v.Product.Description,
			UnitPrice:   float32(v.Product.Price),
			Category:    "marketplace",
			UnitMeasure: "unit",
			Quantity:    uint64(v.Quantity),
			TotalAmount: v.Price,
		})
	}

	return &request.CreatePaymentRequest{
		ExternalReference: externalReference,
		TotalAmount:       order.TotalBill,
		Items:             items,
		Title:             "FIAP Tech Challenge - Product Order",
		Description:       "Purchases made at the FIAP Tech Challenge store",
		NotificationURL:   os.Getenv("MERCADO_PAGO_NOTIFICATION_URL"),
	}
}
