package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewCustomerService, fx.As(new(port.ICustomerService))),
		fx.Annotate(NewProductService, fx.As(new(port.IProductService))),
		fx.Annotate(NewSignInService, fx.As(new(port.ISignInService))),
		fx.Annotate(NewOrderService, fx.As(new(port.IOrderService))),
		fx.Annotate(NewPaymentService, fx.As(new(port.IPaymentService))),
		fx.Annotate(NewOrderHistoryService, fx.As(new(port.IOrderHistoryService))),
		fx.Annotate(NewCategoryService, fx.As(new(port.ICategoryService))),
		fx.Annotate(NewStaffService, fx.As(new(port.IStaffService))),
		fx.Annotate(NewOrderProductService, fx.As(new(port.IOrderProductService))),
	),
)
