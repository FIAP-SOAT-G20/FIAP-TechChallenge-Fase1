package tests

import (
	"time"

	"github.com/go-faker/faker/v4"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

func MockCustomer() *domain.Customer {
	return &domain.Customer{
		Name:      faker.FirstName(),
		Email:     faker.Email(),
		CPF:       "123.456.789-00",
		CreatedAt: time.Now(),
	}
}
