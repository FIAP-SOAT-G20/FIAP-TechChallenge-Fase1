package handler

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

// SignInHandler represents the HTTP handler for authentication-related requests
type SignInHandler struct {
	service port.ISignInService
}

// NewSignInHandler creates a new SignInHandler instance
func NewSignInHandler(service port.ISignInService) *SignInHandler {
	return &SignInHandler{
		service,
	}
}

func (h *SignInHandler) GroupRouterPattern() string {
	return "/api/v1/sign-in"
}

func (h *SignInHandler) Register(router *gin.RouterGroup) {
	router.POST("", h.SignIn)
}

// SignIn godoc
//
//	@Summary		Sign in a customer
//	@Description	Sign in a customer
//	@Description	Example CPF: 123.456.789-00
//	@Description	> 2.b: ii. Identificação do Cliente via CPF
//	@Tags			sign-in
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.SignInRequest	true	"SignIn Request"
//	@Success		200		{object}	response.SignInResponse	"Succesfully signed in"
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		401		{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/sign-in [post]
func (ah *SignInHandler) SignIn(c *gin.Context) {
	var req request.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	customer, err := ah.service.GetByCPF(req.Cpf)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	rsp := response.NewSignInResponse(customer)

	c.JSON(http.StatusOK, rsp)
}
