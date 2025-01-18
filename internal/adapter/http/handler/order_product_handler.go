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

type OrderProductHandler struct {
	service port.IOrderProductService
}

func NewOrderProductHandler(service port.IOrderProductService) *OrderProductHandler {
	return &OrderProductHandler{service: service}
}

func (h *OrderProductHandler) Register(router *gin.RouterGroup) {
	router.GET("/:order_id/:product_id", h.GetOrderProduct)
	router.GET("/", h.ListOrderProducts)
	router.POST("/:order_id/:product_id", h.CreateOrderProduct)
	router.PUT("/:order_id/:product_id", h.UpdateOrderProduct)
	router.DELETE("/:order_id/:product_id", h.DeleteOrderProduct)
}

func (h *OrderProductHandler) GroupRouterPattern() string {
	return "/api/v1/orders/products"
}

// GetOrderProduct godoc
//
//	@Summary		Get an order product
//	@Description	Get an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id		path		int				true	"Order ID"
//	@Param			product_id		path		int				true	"Product ID"
//	@Success		200		{object}	response.OrderProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [get]
func (h *OrderProductHandler) GetOrderProduct(c *gin.Context) {
	var pathParams *request.OrderProductPathParam

	if err := c.ShouldBindUri(&pathParams); err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	orderProduct, err := h.service.GetByID(pathParams.OrderID, pathParams.ProductID)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	orderProductResponse := response.NewOrderProductResponse(orderProduct)
	c.JSON(http.StatusOK, orderProductResponse)
}

// CreateOrderProduct godoc
//
//	@Summary		Create an order product
//	@Description	Create an order product
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Param			order_id		path		int				true	"Order ID"
//	@Param			product_id		path		int				true	"Product ID"
//	@Param			order	body		request.OrderProductRequest	true	"OrderProductRequest"
//	@Success		201		{object}	response.OrderProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [post]
func (h *OrderProductHandler) CreateOrderProduct(c *gin.Context) {
	var pathParams *request.OrderProductPathParam

	if err := c.ShouldBindUri(&pathParams); err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	var req request.OrderProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	orderProduct := &domain.OrderProduct{
		OrderID:   pathParams.OrderID,
		ProductID: pathParams.ProductID,
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
//	@Param			order_id		path		int				true	"Order ID"
//	@Param			product_id		path		int				true	"Product ID"
//	@Param			order	body		request.OrderProductRequest	true	"OrderProductResponse"
//	@Success		201		{object}	response.OrderProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [put]
func (h *OrderProductHandler) UpdateOrderProduct(c *gin.Context) {
	var pathParams *request.OrderProductPathParam

	if err := c.ShouldBindUri(&pathParams); err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	var req request.OrderProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	orderProduct := &domain.OrderProduct{
		OrderID:   pathParams.OrderID,
		ProductID: pathParams.ProductID,
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
//	@Param			order_id		path		int				true	"Order ID"
//	@Param			product_id		path		int				true	"Product ID"
//	@Success		204	{object}	string
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/products/{order_id}/{product_id} [delete]
func (h *OrderProductHandler) DeleteOrderProduct(c *gin.Context) {
	var pathParams *request.OrderProductPathParam

	if err := c.ShouldBindUri(&pathParams); err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	if err := h.service.Delete(pathParams.OrderID, pathParams.ProductID); err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
