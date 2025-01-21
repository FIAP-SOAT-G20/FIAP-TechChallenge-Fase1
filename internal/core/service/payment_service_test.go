package service

import (
	"reflect"
	"testing"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/mock/repository"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPaymentService(t *testing.T) {
	type args struct {
		paymentRepository        port.IPaymentRepository
		orderRepository          port.IOrderRepository
		paymentGatewayReposiroty port.IPaymentGatewayRepository
	}
	tests := []struct {
		name string
		args args
		want *PaymentService
	}{
		{
			name: "TestNewPaymentService",
			args: args{
				paymentRepository:        nil,
				orderRepository:          nil,
				paymentGatewayReposiroty: nil,
			},
			want: &PaymentService{
				paymentRepository:        nil,
				orderRepository:          nil,
				paymentGatewayRepository: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPaymentService(tt.args.paymentRepository, tt.args.orderRepository, tt.args.paymentGatewayReposiroty); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaymentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_CreatePayment(t *testing.T) {
	mockPaymentRepository := new(mocks.MockPaymentRepository)
	mockOrderRepository := new(repository.OrderRepositoryMock)
	mockPaymentGatewayRepository := new(mocks.MockPaymentGatewayRepository)

	type args struct {
		orderID uint64
	}
	tests := []struct {
		name       string
		args       args
		setupMocks func()
		want       *domain.Payment
		wantErr    bool
	}{
		{
			name: "Success - Existing Pending Payment",
			args: args{
				orderID: 0,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(0)).
					Return(&domain.Payment{
						ID: 1,
					}, nil)
			},
			want: &domain.Payment{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "Success - New Payment",
			args: args{
				orderID: 1,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(1)).
					Return(&domain.Payment{
						ID: 0,
					}, nil)

				mockOrderRepository.
					On("GetByID", uint64(1)).
					Return(&domain.Order{ID: 1}, nil)

				mockPaymentGatewayRepository.
					On("CreatePayment", &domain.CreatePaymentIN{
						ExternalReference: "1",
						TotalAmount:       0,
						Items:             nil,
						Title:             "FIAP Tech Challenge - Product Order",
						Description:       "Purchases made at the FIAP Tech Challenge store",
						NotificationUrl:   "",
					}).
					Return(&domain.CreatePaymentOUT{
						InStoreOrderID: "123",
						QrData:         "456",
					}, nil)

				mockPaymentRepository.
					On("Insert", &domain.Payment{
						Status:            domain.PROCESSING,
						ExternalPaymentID: "123",
						OrderID:           uint64(1),
						QrData:            "456",
					}).
					Return(&domain.Payment{ID: 1}, nil)
			},
			want: &domain.Payment{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name: "Error - GetPaymentByOrderIDAndStatus",
			args: args{
				orderID: 0,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(0)).
					Return(nil, assert.AnError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - GetByID",
			args: args{
				orderID: 1,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(1)).
					Return(&domain.Payment{
						ID: 0,
					}, nil)

				mockOrderRepository.
					On("GetByID", uint64(1)).
					Return(nil, assert.AnError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - CreatePaymentMock",
			args: args{
				orderID: 1,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(1)).
					Return(&domain.Payment{
						ID: 0,
					}, nil)

				mockOrderRepository.
					On("GetByID", uint64(1)).
					Return(&domain.Order{ID: 1}, nil)

				mockPaymentGatewayRepository.
					On("CreatePaymentMock", &domain.CreatePaymentIN{
						ExternalReference: "1",
						TotalAmount:       0,
						Items:             nil,
						Title:             "FIAP Tech Challenge - Product Order",
						Description:       "Purchases made at the FIAP Tech Challenge store",
						NotificationUrl:   "",
					}).
					Return(nil, assert.AnError)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - Insert - Payment",
			args: args{
				orderID: 1,
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("GetPaymentByOrderIDAndStatus", domain.PROCESSING, uint64(1)).
					Return(&domain.Payment{ID: 0}, nil)

				mockPaymentRepository.
					On("Insert", mock.Anything).
					Return(nil, assert.AnError)

				mockOrderRepository.
					On("GetByID", mock.Anything).
					Return(&domain.Order{ID: 1}, nil)

				mockPaymentGatewayRepository.
					On("CreatePaymentMock", mock.Anything).
					Return(&domain.CreatePaymentOUT{
						InStoreOrderID: "123",
						QrData:         "456",
					}, nil)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			t.Cleanup(func() {
				mockPaymentGatewayRepository.ExpectedCalls = nil
				mockOrderRepository.ExpectedCalls = nil
				mockPaymentRepository.ExpectedCalls = nil
			})
			tt.setupMocks()
			ps := NewPaymentService(
				mockPaymentRepository,
				mockOrderRepository,
				mockPaymentGatewayRepository,
			)

			// Act
			got, err := ps.CreatePayment(tt.args.orderID)

			// Assert
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPaymentService_createPaymentPayload(t *testing.T) {
	type args struct {
		order *domain.Order
	}
	tests := []struct {
		name string
		args args
		want *domain.CreatePaymentIN
	}{
		{
			name: "Success",
			args: args{
				order: &domain.Order{
					ID: 1,
					OrderProducts: []domain.OrderProduct{
						{
							ProductID: 1,
							Price:     10,
							Quantity:  1,
						},
					},
				},
			},
			want: &domain.CreatePaymentIN{
				ExternalReference: "1",
				TotalAmount:       0,
				Items: []domain.ItemsIN{
					{
						Title:       "",
						Description: "",
						UnitPrice:   0,
						Category:    "marketplace",
						UnitMeasure: "unit",
						Quantity:    1,
						TotalAmount: 10,
					},
				},
				Title:           "FIAP Tech Challenge - Product Order",
				Description:     "Purchases made at the FIAP Tech Challenge store",
				NotificationUrl: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			if got := createPaymentGatewayPayload(tt.args.order); !reflect.DeepEqual(got, tt.want) {

				// Assert
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPaymentService_UpdatePayment(t *testing.T) {
	mockPaymentRepository := new(mocks.MockPaymentRepository)
	mockOrderRepository := new(repository.OrderRepositoryMock)

	type args struct {
		payment *domain.UpdatePaymentIN
	}
	tests := []struct {
		name       string
		args       args
		setupMocks func()
		want       *domain.Payment
		wantErr    bool
	}{
		{
			name: "Success",
			args: args{
				payment: &domain.UpdatePaymentIN{
					Resource: "123",
					Topic:    "topic",
				},
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("UpdateStatus", domain.CONFIRMED, "123").
					Return(nil)

				mockPaymentRepository.
					On("GetByExternalPaymentID", "123").
					Return(&domain.Payment{
						ID: 1,
					}, nil)

				mockOrderRepository.
					On("GetByID", mock.Anything).
					Return(&domain.Order{ID: 1}, nil)

				mockOrderRepository.
					On("UpdateStatus", mock.Anything).
					Return(nil)
			},
			want:    &domain.Payment{ID: 1},
			wantErr: false,
		},
		{
			name: "Error - UpdateStatus",
			args: args{
				payment: &domain.UpdatePaymentIN{
					Resource: "123",
					Topic:    "topic",
				},
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("UpdateStatus", domain.CONFIRMED, "123").
					Return(assert.AnError)

				mockOrderRepository.
					On("GetByID", mock.Anything).
					Return(&domain.Order{ID: 1}, nil)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error - GetByExternalPaymentID",
			args: args{
				payment: &domain.UpdatePaymentIN{
					Resource: "123",
					Topic:    "topic",
				},
			},
			setupMocks: func() {
				mockPaymentRepository.
					On("UpdateStatus", domain.CONFIRMED, "123").
					Return(nil)

				mockPaymentRepository.
					On("GetByExternalPaymentID", "123").
					Return(nil, assert.AnError)

				mockOrderRepository.
					On("GetByID", mock.Anything).
					Return(&domain.Order{ID: 1}, nil)

			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			t.Cleanup(func() {
				mockPaymentRepository.ExpectedCalls = nil
			})
			tt.setupMocks()
			ps := NewPaymentService(mockPaymentRepository, mockOrderRepository, nil)

			// Act
			got, err := ps.UpdatePayment(tt.args.payment)

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentService.UpdatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentService.UpdatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
