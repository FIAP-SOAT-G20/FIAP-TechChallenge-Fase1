package server

import (
	"context"
	"fmt"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/router"
	"go.uber.org/fx"
	"log/slog"
)

func StartServer(lifecycle fx.Lifecycle, router *router.Router, environment *config.Environment) {
	listenAddress := fmt.Sprintf(":%s", environment.Port)
	slog.Info("Starting the HTTP server", "address", listenAddress)
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err := router.Serve(listenAddress)
			if err != nil {
				slog.Error("Failed to start server", "error", err)
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Stopping the HTTP server")
			return nil
		},
	})
}
