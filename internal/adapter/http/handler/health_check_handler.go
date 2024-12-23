package handler

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Register(router *gin.RouterGroup) {
	router.GET("", h.HealthCheck)
	router.GET("/", h.HealthCheck)
}

func (h *HealthCheckHandler) GroupRouterPattern() string {
	return "/healthCheck"
}

// HealthCheck godoc
//
//	@Summary		Application HealthCheck
//	@Description	Checks application health
//	@Tags			healthCheck
//	@Produce		json
//	@Success		200			{object}	response.HealthCheckResponse
//	@Failure		500			{object}	response.ErrorResponse	"Internal server error"
//	@Router			/healthCheck [GET]
func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response.NewHealthCheckResponse())
}
