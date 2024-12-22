package internal

import (
	adapters "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/server"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core"
	"go.uber.org/fx"
)

func NewApp(cfg *config.Environment) *fx.App {
	return fx.New(
		fx.Provide(func() *config.Environment { return cfg }),
		core.Module,
		adapters.Module,
		fx.Invoke(server.StartServer),
	)
}
