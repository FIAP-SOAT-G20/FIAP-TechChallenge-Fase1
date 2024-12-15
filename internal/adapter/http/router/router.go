package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/docs"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/handler"
)

type Router struct {
	productHandler  *handler.ProductHandler
	customerHandler *handler.CustomerHandler
}

func NewRouter(productHandler *handler.ProductHandler, customerHandler *handler.CustomerHandler) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	docs.SwaggerInfo.BasePath = ""
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productHandler.Register(router.Group("/api/v1/products"))
	customerHandler.Register(router.Group("/api/v1/customers"))

	return router
}
