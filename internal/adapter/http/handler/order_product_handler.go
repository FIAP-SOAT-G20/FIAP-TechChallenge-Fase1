package handler

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderProductHandler struct {
	service port.IOrderProductService
}

func NewOrderProductHandler(service port.IOrderProductService) *OrderProductHandler {
	return &OrderProductHandler{service: service}
}

func (h *OrderProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListOrderProducts)
	router.POST("/", h.CreateOrderProduct)
	router.PUT("/", h.UpdateOrderProduct)
	router.DELETE("/", h.DeleteOrderProduct)
}

func (h *OrderProductHandler) GroupRouterPattern() string {
	return "/api/v1/orders/products"
}

// CreateOrderProduct godoc
//
//	@Summary		Create an order product
//	@Description	Create an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		request.CreateOrderProductRequest	true	"CreateOrderProductRequest"
//	@Success		201		{object}	response.OrderProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products [post]
func (h *OrderProductHandler) CreateOrderProduct(c *gin.Context) {
	var req request.CreateOrderProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	orderProduct := &domain.OrderProduct{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.service.Create(orderProduct); err != nil {
		response.HandleError(c, err)
		return
	}

	orderProductResponse := response.NewOrderProductResponse(orderProduct)
	c.JSON(http.StatusCreated, orderProductResponse)
}

// UpdateOrderProduct godoc
//
//	@Summary		Update an order product
//	@Description	Update an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order	body		request.UpdateOrderProductRequest	true	"OrderProductResponse"
//	@Success		201		{object}	response.OrderProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products [put]
func (h *OrderProductHandler) UpdateOrderProduct(c *gin.Context) {
	var req request.UpdateOrderProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	orderProduct := &domain.OrderProduct{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.service.Update(orderProduct); err != nil {
		response.HandleError(c, err)
		return
	}

	orderProductResponse := response.NewOrderProductResponse(orderProduct)
	c.JSON(http.StatusOK, orderProductResponse)
}

// ListOrderProducts godoc
//
//	@Summary		List order products
//	@Description	List order products
//	@Tags			orderHistories
//	@Accept			json
//	@Produce		json
//	@Param			order_id	query		uint64	false	"Order ID"
//	@Param			product_id	query		uint64	false	"Product ID"
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Success		200			{object}	response.OrderProductPaginated
//	@Router			/api/v1/orders/products [get]
func (h *OrderProductHandler) ListOrderProducts(c *gin.Context) {
	orderID := c.DefaultQuery("order_id", "0")
	productID := c.DefaultQuery("product_id", "0")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	orderIDUint64, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidQueryParams)
		return
	}

	productIDUint64, err := strconv.ParseUint(productID, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidQueryParams)
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidQueryParams)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidQueryParams)
		return
	}

	orders, total, err := h.service.List(orderIDUint64, productIDUint64, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	responses := response.NewOrderProductPaginated(orders, total, pageInt, limitInt)
	c.JSON(http.StatusOK, responses)
}

// DeleteOrderProduct godoc
//
//	@Summary		Delete an order product
//	@Description	Delete an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"OrderResponse ID"
//	@Param			orderProduct	body		request.DeleteOrderProductRequest	true	"DeleteOrderProductRequest"
//	@Success		204	{object}	string
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products/ [delete]
func (h *OrderProductHandler) DeleteOrderProduct(c *gin.Context) {
	var req request.DeleteOrderProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	orderProduct := &domain.OrderProduct{
		OrderID:   req.OrderID,
		ProductID: req.ProductID,
	}

	if err := h.service.Delete(orderProduct); err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
