package router

import (
	"log/slog"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/docs"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/handler"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	environment string, 
	signInHandler *handler.SignInHandler, 
	productHandler *handler.ProductHandler, 
	customerHandler *handler.CustomerHandler,
) *Router {
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(sloggin.New(slog.Default()), gin.Recovery())

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

	signInHandler.Register(router.Group("/api/v1/sign-in"))
	productHandler.Register(router.Group("/api/v1/products"))
	customerHandler.Register(router.Group("/api/v1/customers"))

	return &Router{Engine: router}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Engine.Run(listenAddr)
}
