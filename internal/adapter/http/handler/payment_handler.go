package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type PaymentHandler struct {
	service port.IPaymentService
}

func NewPaymentHandler(paymentService port.IPaymentService) *PaymentHandler {
	return &PaymentHandler{service: paymentService}
}

func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	router.POST("/:order_id/checkout", h.CreatePayment)
	router.POST("/callback", h.UpdatePayment)
}

func (h *PaymentHandler) GroupRouterPattern() string {
	return "/api/v1/payments"
}

// CreatePayment godoc
//
//	@Summary		Create a checkout on a order
//	@Description	Create a checkout on a order (2.b: > 2.b.: v. Fake checkout)
//	@Tags			products, payments
//	@Accept			json
//	@Produce		json
//	@Param			order_id		path		int				true	"Order ID"
//	@Success		200		{object}	response.PaymentResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/payments/{order_id}/checkout [post]
func (h PaymentHandler) CreatePayment(c *gin.Context) {
    var pathParams request.PaymentPathParam

    if err := c.ShouldBindUri(&pathParams); err != nil {
        response.HandleError(c, domain.ErrInvalidParam)
        return
    }

    orderIDUint64, err := strconv.ParseUint(pathParams.OrderID, 10, 64)
    if err != nil {
        response.HandleError(c, domain.ErrInvalidParam)
        return
    }

    payment, err := h.service.CreatePayment(orderIDUint64)
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
//	@Description	Update a payment on a order (2.b: > 2.b.: v. Fake checkout)
// 	@Description	- resource = external payment id, obtained from the checkout response
// 	@Description	- topic = payment
//	@Tags			products, payments
//	@Accept			json
//	@Produce		json
//	@Param			product	body		request.UpdatePaymentRequest	true	"Update Payment Request"
//	@Success		200		{object}	response.PaymentResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/payments/callback [post]
func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
	var req request.UpdatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	paymentIN := req.ToDomain()

	payment, err := h.service.UpdatePayment(paymentIN)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	paymentReponse := response.NewPaymentResponse(payment)
	c.JSON(http.StatusOK, paymentReponse)
}
