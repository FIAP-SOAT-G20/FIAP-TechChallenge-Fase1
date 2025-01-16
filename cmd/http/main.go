package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/logger"
)

//	@title			FIAP Tech Challenge Fase 1 - 10SOAT - G20
//	@version		1
//	@description	API de um Fast Food para o Tech Challenge da FIAP - Fase 1 - 10SOAT - G20
//	@servers		[ { "url": "http://localhost:8080" } ]
//	@host			localhost:8080
//	@BasePath		/api/v1
//	@tag.name sign-up
//	@tag.description 2.b: i.Cadastro do Cliente
//	@tag.name products
//	@tag.description 2.b: iii. Criar, editar e remover produtos;
//	@tag.name payments
//	@tag.description 2.b: v. Fake checkout
//	@tag.name sign-in
//	@tag.description 2.b: ii. Identificação do Cliente via CPF
// @externalDocs.description GitHub Repository
// @externalDocs.url https://github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1
//
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
