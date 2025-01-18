package handler

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		AsRoute(NewCustomerHandler),
		AsRoute(NewOrderProductHandler),
		AsRoute(NewOrderHistoryHandler),
		AsRoute(NewProductHandler),
		AsRoute(NewHealthCheckHandler),
		AsRoute(NewSignInHandler),
		AsRoute(NewOrderHandler),
		AsRoute(NewPaymentHandler),
		AsRoute(NewCategoryHandler),
		AsRoute(NewStaffHandler),
	),
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(router.IRouter)),
		fx.ResultTags(`group:"routes"`),
	)
}
