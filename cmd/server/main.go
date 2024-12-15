package main

import (
	"log"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/handler"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/router"
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

	database, err := postgres.Connect(environment.DatabaseURL)
	if err != nil {
		log.Fatalln("failed to connect database", err)
	}

	// repositories
	categoryRepository := repository.NewCategoryRepository(database)
	productRepository := repository.NewProductRepository(database)
	customerRepository := repository.NewCustomerRepository(database)

	// services
	productServive := service.NewProductService(productRepository, categoryRepository)
	customerService := service.NewCustomerService(customerRepository)

	// handlers
	productHandler := handler.NewProductHandler(productServive)
	customerHandler := handler.NewCustomerHandler(customerService)

	// router
	routes := router.NewRouter(productHandler, customerHandler)
	if err := routes.Run(":" + environment.Port); err != nil {
		log.Fatalln("failed to run server", err)
	}
}
