package repository

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(fx.Annotate(NewCustomerRepository, fx.As(new(port.ICustomerRepository)))),
	fx.Provide(fx.Annotate(NewCategoryRepository, fx.As(new(port.ICategoryRepository)))),
	fx.Provide(fx.Annotate(NewProductRepository, fx.As(new(port.IProductRepository)))),
	fx.Provide(fx.Annotate(NewOrderRepository, fx.As(new(port.IOrderRepository)))),
	fx.Provide(fx.Annotate(NewOrderHistoryRepository, fx.As(new(port.IOrderHistoryRepository)))),
	fx.Provide(fx.Annotate(NewOrderProductRepository, fx.As(new(port.IOrderProductRepository)))),
)
