package repository

import (
	"fmt"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/go-resty/resty/v2"
)

type PaymentGatewayRepository struct {
	cfg *config.Environment
}

func NewPaymentGatewayRepository(cfg *config.Environment) *PaymentGatewayRepository {
	return &PaymentGatewayRepository{
		cfg: cfg,
	}
}

func (ps *PaymentGatewayRepository) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	payment.NotificationUrl = ps.cfg.PaymentGatewayNotificationURL
	body := request.NewPaymentRequest(payment)

	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", ps.cfg.PaymentGatewayToken)).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(body).
		SetResult(&response.CreatePaymentResponse{}).
		Post(ps.cfg.PaymentGatewayURL)
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	response := response.ToCreatePaymentOUTDomain(resp.Result().(*response.CreatePaymentResponse))

	return response, nil
}
