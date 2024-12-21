package main

import (
	"context"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/logger"
	"log"
	"log/slog"
)

//	@title			FIAP Tech Challenge Fase 1 - G20 - 10 SOAT
//	@version		1
//	@description	API para o Tech Challenge da FIAP - Fase 1 - G20 - 10 SOAT

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	environment, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalln(err)
	}

	// logger
	logger.Set(environment.AppEnvironment)
	slog.Info("Starting the application", "app", "TC 01 G20 10SOAT", "env", environment.AppEnvironment)

	application := internal.NewApp(environment)
	err = application.Start(context.Background())
	if err != nil {
		slog.Error("Error starting FX app", "env", environment.AppEnvironment)
		return
	}
}
