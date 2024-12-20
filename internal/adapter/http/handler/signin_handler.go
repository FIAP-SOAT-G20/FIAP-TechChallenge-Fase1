package handler

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

// SignInHandler represents the HTTP handler for authentication-related requests
type SignInHandler struct {
	svc port.ISignInService
}

// NewSignInHandler creates a new SignInHandler instance
func NewSignInHandler(svc port.ISignInService) *SignInHandler {
	return &SignInHandler{
		svc,
	}
}

func (h *SignInHandler) GroupRouterPattern() string {
	return "/api/v1/sign-in"
}

func (h *SignInHandler) Register(router *gin.RouterGroup) {
	router.POST("", h.SignIn)
}

// signInRequest represents the request body for logging in a user
type signInRequest struct {
	Cpf string `json:"cpf" binding:"required" example:"000.000.000-00"`
}

// SignIn godoc
//
//	@Summary		Sign in a customer
//	@Description	Sign in a customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			request	body		signInRequest			true	"SignInResponse"
//	@Success		200		{object}	response.SignInResponse	"Succesfully signed in"
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		401		{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/sign-in [post]
func (ah *SignInHandler) SignIn(c *gin.Context) {
	var req signInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	customer, err := ah.svc.GetByCPF(req.Cpf)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	rsp := response.NewSignInResponse(customer)

	c.JSON(http.StatusOK, rsp)
}
