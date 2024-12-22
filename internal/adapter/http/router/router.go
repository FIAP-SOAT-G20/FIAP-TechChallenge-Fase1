package router

import (
	"log/slog"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/docs"
)

type IRouter interface {
	GroupRouterPattern() string
	Register(router *gin.RouterGroup)
}

type Router struct {
	Engine *gin.Engine
}

func InitGinEngine(cfg *config.Environment) {
	if cfg.AppEnvironment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func NewRouter(handlers []IRouter) *Router {
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

	for _, handler := range handlers {
		handler.Register(router.Group(handler.GroupRouterPattern()))
	}

	return &Router{Engine: router}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Engine.Run(listenAddr)
}
