package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PathParam struct {
	OrderID uint64 `uri:"orderId" binding:"required"`
}

type PaymentHandler struct {
	paymentService port.IPaymentService
}

func NewPaymentHandler(paymentService port.IPaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	router.POST("/:orderId/checkout", h.CreatePayment)
	router.POST("/callback", h.UpdatePayment)
}

func (h *PaymentHandler) GroupRouterPattern() string {
	return "/api/v1/payments"
}

// CreatePayment godoc
//
//	@Summary		Create a checkout on a order
//	@Description	Create a checkout on a order
//	@Tags			products, payments
//	@Accept			json
//	@Produce		json
//	@Param			orderId		path		int				true	"Order ID"
//	@Success		200		{object}	response.PaymentResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/payments/{orderId}/checkout [put]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var pathParams *PathParam

	if err := c.ShouldBindUri(&pathParams); err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	payment, err := h.paymentService.CreatePayment(pathParams.OrderID)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	paymentReponse := response.NewPaymentResponse(payment)
	c.JSON(http.StatusCreated, paymentReponse)
}

// Update Payment godoc
//
//	@Summary		Update a payment on a order
//	@Description	Update a payment on a order
//	@Tags			products, payments
//	@Accept			json
//	@Produce		json
//	@Param			product	body		request.UpdatePaymentRequest	true	"PaymentResponse"
//	@Success		200		{object}	response.PaymentResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/callback [put]
func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	var req request.UpdatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	paymentIN := request.NewUpdatePaymentRequest(&req)

	payment, err := h.paymentService.UpdatePayment(paymentIN)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	paymentReponse := response.NewPaymentResponse(payment)
	c.JSON(http.StatusOK, paymentReponse)
}
