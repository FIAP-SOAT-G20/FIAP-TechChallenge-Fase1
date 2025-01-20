package repository

import (
	"github.com/google/uuid"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type FakePaymentGatewayRepository struct {
	cfg *config.Environment
}

func NewFakePaymentGatewayRepository(cfg *config.Environment) *FakePaymentGatewayRepository {
	return &FakePaymentGatewayRepository{
		cfg: cfg,
	}
}

func (ps *FakePaymentGatewayRepository) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	return &domain.CreatePaymentOUT{
		InStoreOrderID: uuid.New().String(),
		QrData:         "https://www.fiap-10-soat-g20.com.br/qr/123456",
	}, nil
}
