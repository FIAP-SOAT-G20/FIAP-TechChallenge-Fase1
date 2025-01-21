package mocks

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockPaymentGatewayRepository struct {
	mock.Mock
}

func (m *MockPaymentGatewayRepository) CreatePayment(payment *domain.CreatePaymentIN) (*domain.CreatePaymentOUT, error) {
	ret := m.Called(payment)

	var r0 *domain.CreatePaymentOUT
	if rf, ok := ret.Get(0).(func(*domain.CreatePaymentIN) *domain.CreatePaymentOUT); ok {
		r0 = rf(payment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CreatePaymentOUT)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.CreatePaymentIN) error); ok {
		r1 = rf(payment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
