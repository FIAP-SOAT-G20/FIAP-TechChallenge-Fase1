package handler

import (
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderHandler struct {
	service port.IOrderService
}

func NewOrderHandler(service port.IOrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateOrder)
	router.GET("/", h.ListOrders)
	router.GET("/:id", h.GetOrder)
	router.PUT("/:id", h.UpdateOrder)
	router.DELETE("/:id", h.DeleteOrder)
}

func (h *OrderHandler) GroupRouterPattern() string {
	return "/api/v1/orders"
}

// CreateOrder godoc
//
//	@Summary		Create an order
//	@Description	Create an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			createOrderRequest	body		request.CreateOrderRequest	true	"Create Order Request"
//	@Success		201		{object}	response.OrderResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	order := &domain.Order{
		CustomerID: req.CustomerID,
	}

	if err := h.service.Create(order); err != nil {
		response.HandleError(c, err)
		return
	}

	orderResponse := response.NewOrderResponse(order)
	c.JSON(http.StatusCreated, orderResponse)
}

// GetOrder godoc
//
//	@Summary		Get an order
//	@Description	Get an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		200	{object}	response.OrderResponse
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	order, err := h.service.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	orderResponse := response.NewOrderResponse(order)
	c.JSON(http.StatusOK, orderResponse)
}

// ListOrders godoc
//
//	@Summary		List orders
//	@Description	List orders
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			customer_id	query		int		false	"Customer ID"
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Success		200			{object}	response.OrderPaginated
//	@Router			/api/v1/orders [get]
func (h *OrderHandler) ListOrders(c *gin.Context) {
	customerID := c.DefaultQuery("customer_id", "0")
	status := c.DefaultQuery("status", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	customerIDUint64, err := strconv.ParseUint(customerID, 10, 64)
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

	orderStatus := domain.OrderStatus(status)
	orders, total, err := h.service.List(customerIDUint64, &orderStatus, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	responses := response.NewOrderPaginated(orders, total, pageInt, limitInt)
	c.JSON(http.StatusOK, responses)
}

// UpdateOrder godoc
//
//	@Summary		Update an order
//	@Description	Update an order
// 	@Description	Only staff can update the order status
// 	@Description	The status are: OPEN, CANCELLED, PENDING, RECEIVED, PREPARING, READY, COMPLETED
// 	@Description	Transition of status: 
// 	@Description	- OPEN      -> CANCELLED || PENDING
// 	@Description	- CANCELLED -> {},
// 	@Description	- PENDING   -> OPEN || RECEIVED
// 	@Description	- RECEIVED  -> PREPARING
// 	@Description	- PREPARING -> READY
// 	@Description	- READY     -> COMPLETED
// 	@Description	- COMPLETED -> {}
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Order ID"
//	@Param			updateOrderRequest	body		request.UpdateOrderRequest	true	"Update Order Request"
//	@Success		200		{object}	response.OrderResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	var req request.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	order, err := h.service.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}
	order.Status = req.Status

	if err := h.service.Update(order, req.StaffID); err != nil {
		response.HandleError(c, err)
		return
	}

	orderResponse := response.NewOrderResponse(order)
	c.JSON(http.StatusOK, orderResponse)
}

// DeleteOrder godoc
//
//	@Summary		Delete an order
//	@Description	Delete an order
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order ID"
//	@Success		204	{object}	string
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	if err := h.service.Delete(idUint64); err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
