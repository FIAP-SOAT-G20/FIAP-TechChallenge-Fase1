package handler

import (
	"net/http"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

// SignInHandler represents the HTTP handler for authentication-related requests
type SignInHandler struct {
	svc port.SigninService
}

// NewSignInHandler creates a new SignInHandler instance
func NewSignInHandler(svc port.SigninService) *SignInHandler {
	return &SignInHandler{
		svc,
	}
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
//	@Description 	Sign in a customer
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		signInRequest	true	"Sign in request body"
//	@Success		200		{object}	authResponse	"Succesfully signed in"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		401		{object}	errorResponse	"Unauthorized error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/sign-in [post]
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
