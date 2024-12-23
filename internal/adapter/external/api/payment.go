package api

import (
	"fmt"
	"os"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type ExternalPaymentService struct {
}

func NewExternalPaymentService() *ExternalPaymentService {
	return &ExternalPaymentService{}
}

func (ps *ExternalPaymentService) CreatePayment(payment *request.CreatePaymentRequest) (*response.CreatePaymentResponse, error) {
	client := resty.New().
		SetTimeout(10*time.Second).
		SetRetryCount(2).
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("MERCADO_PAGO_TOKEN"))).
		SetHeader("Content-Type", "application/json")

	resp, err := client.R().
		SetBody(payment).
		SetResult(&response.CreatePaymentResponse{}).
		Post(os.Getenv("MERCADO_PAGO_URL"))
	if err != nil {
		return nil, fmt.Errorf("error to create payment: %w", err)
	}

	if resp.StatusCode() != 201 {
		logrus.Infof("Response body: %s", resp.String())
		return nil, fmt.Errorf("error: response status %d", resp.StatusCode())
	}

	return resp.Result().(*response.CreatePaymentResponse), nil
}
