package api

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/go-resty/resty/v2"
)

type ExternalPaymentService struct {
}

func NewExternalPaymentService() *ExternalPaymentService {
	return &ExternalPaymentService{}
}

func (ps *ExternalPaymentService) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	body := request.NewPaymentRequest(payment)

	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("MERCADO_PAGO_TOKEN"))).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(body).
		SetResult(&response.CreatePaymentResponse{}).
		Post(os.Getenv("MERCADO_PAGO_URL"))
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	response := response.ToCreatePaymentOUTDomain(resp.Result().(*response.CreatePaymentResponse))

	return response, nil
}

func (ps *ExternalPaymentService) CreatePaymentMock(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	return &domain.CreatePaymentOUT{
		InStoreOrderID: uuid.New().String(),
		QrData:         "https://www.fiap-10-soat-g20.com.br/qr/123456",
	}, nil
}
