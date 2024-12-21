package router

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(InitGinEngine),
	fx.Provide(fx.Annotate(NewRouter, fx.ParamTags(`group:"routes"`))),
)
