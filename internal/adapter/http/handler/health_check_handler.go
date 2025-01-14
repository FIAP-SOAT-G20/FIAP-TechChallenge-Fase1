package handler

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/gin-gonic/gin"
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
	return "/health"
}

// HealthCheck godoc
//
//	@Summary		Application HealthCheck
//	@Description	Checks application health
//	@Tags			healthcheck
//	@Produce		json
//	@Success		200			{object}	response.HealthCheckResponse
//	@Failure		500			{object}	response.ErrorResponse "Internal server error"
//	@Failure		503			{object}	response.HealthCheckResponse "Service Unavailable"
//	@Router			/health [GET]
func (h *HealthCheckHandler) HealthCheck(c *gin.Context) {
	hc := &response.HealthCheckResponse{
		Status: response.HealthCheckStatusPass,
		Checks: map[string]response.HealthCheckVerifications{
			"db:postgres": {
				ComponentId: "db:postgres",
				Status:      response.HealthCheckStatusPass,
				Time:        time.Now(),
			},
		},
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		response.HandleError(c, err)
		return
	}

	if db.Ping() != nil {
		hc.Status = response.HealthCheckStatusFail
		hc.Checks["db:postgres"] = response.HealthCheckVerifications{
			ComponentId: "db:postgres",
			Status:      response.HealthCheckStatusFail,
			Time:        time.Now(),
		}
		c.JSON(http.StatusServiceUnavailable, hc)
		return
	}

	switch hc.Status {
	case response.HealthCheckStatusFail:
		c.JSON(http.StatusServiceUnavailable, hc)
	default:
		c.JSON(http.StatusOK, hc)
	}
}
