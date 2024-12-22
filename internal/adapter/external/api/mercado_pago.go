package api

import (
	"fmt"
	"os"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type MercadoPagoService struct {
}

func NewMercadoPagoService() *MercadoPagoService {
	return &MercadoPagoService{}
}

func (ps *MercadoPagoService) CreatePayment(payment *domain.CreatePayment) (*domain.CreatePaymentResponse, error) {
	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("MERCADO_PAGO_TOKEN"))).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(payment).
		SetResult(&domain.CreatePaymentResponse{}).
		Post("https://api.mercadopago.com/instore/orders/qr/seller/collectors/339709477/pos/FTC01/qrs")
	if err != nil {
		return nil, fmt.Errorf("error to create mercado pago payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		logrus.Infof("Response body: %s", resp.String())
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	return resp.Result().(*domain.CreatePaymentResponse), nil
}
