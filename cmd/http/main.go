package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/router"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/logger"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/postgres"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/postgres/repository"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/service"
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

	// init database connection
	dbConnection, err := postgres.New(environment.DatabaseURL)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	// migrate database
	if err = dbConnection.Migrate(); err != nil {
		slog.Error("error migrating database", "error", err)
		os.Exit(1)
	}

	// repositories
	categoryRepository := repository.NewCategoryRepository(dbConnection.DB)
	productRepository := repository.NewProductRepository(dbConnection.DB)
	customerRepository := repository.NewCustomerRepository(dbConnection.DB)

	// services
	signInService := service.NewSignInService(customerRepository)
	productServive := service.NewProductService(productRepository, categoryRepository)
	customerService := service.NewCustomerService(customerRepository)

	// handlers
	signInHandler := handler.NewSignInHandler(signInService)
	productHandler := handler.NewProductHandler(productServive)
	customerHandler := handler.NewCustomerHandler(customerService)

	// router
	listenAddress := fmt.Sprintf(":%s", environment.Port)
	slog.Info("Starting the HTTP server", "address", listenAddress)

	routes := router.NewRouter(
		environment.AppEnvironment,
		signInHandler,
		productHandler,
		customerHandler,
	)

	if err := routes.Serve(listenAddress); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
