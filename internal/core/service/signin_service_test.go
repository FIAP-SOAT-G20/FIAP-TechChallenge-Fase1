package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/tests"
)

func TestSignInService_GetByCPF(t *testing.T) {
	mockCustomerRepository := new(tests.MockCustomerRepository)
	signInService := NewSignInService(mockCustomerRepository)

	scenarios := []struct {
		name          string
		cpf           string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Given valid CPF When GetByCPF is called Then should succeed",
			cpf:  "12345678900",
			setupMocks: func() {
				mockCustomerRepository.On("GetByCPF", "12345678900").Return(tests.MockCustomer(), nil)
			},
			expectedError: nil,
		},
		{
			name: "Given invalid CPF When GetByCPF is called Then should return error",
			cpf:  "12345678900",
			setupMocks: func() {
				mockCustomerRepository.On("GetByCPF", "12345678900").Return(nil, domain.ErrNotFound)
			},
			expectedError: domain.ErrNotFound,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				mockCustomerRepository.ExpectedCalls = nil
			})

			tt.setupMocks()
			_, err := signInService.GetByCPF(tt.cpf)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}
