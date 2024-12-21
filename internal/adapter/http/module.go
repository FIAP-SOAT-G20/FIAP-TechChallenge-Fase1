package http

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/router"
	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	router.Module,
)
