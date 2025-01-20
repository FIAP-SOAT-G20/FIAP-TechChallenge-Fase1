package repository

import (
	"github.com/google/uuid"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type FakeFakePaymentGatewayRepository struct {
	cfg *config.Environment
}

func NewFakePaymentGatewayRepository(cfg *config.Environment) *FakeFakePaymentGatewayRepository {
	return &FakeFakePaymentGatewayRepository{
		cfg: cfg,
	}
}

func (ps *FakeFakePaymentGatewayRepository) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	return &domain.CreatePaymentOUT{
		InStoreOrderID: uuid.New().String(),
		QrData:         "https://www.fiap-10-soat-g20.com.br/qr/123456",
	}, nil
}
