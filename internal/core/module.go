package core

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	service.Module,
)
