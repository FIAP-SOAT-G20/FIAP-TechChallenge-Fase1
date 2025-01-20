package repository

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// fx.Annotate(NewPaymentGatewayRepository, fx.As(new(port.IPaymentGatewayRepository))),
		fx.Annotate(NewFakePaymentGatewayRepository, fx.As(new(port.IPaymentGatewayRepository))),
	),
)
