package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type OrderHistoryHandler struct {
	service port.IOrderHistoryService
}

func NewOrderHistoryHandler(service port.IOrderHistoryService) *OrderHistoryHandler {
	return &OrderHistoryHandler{service: service}
}

func (h *OrderHistoryHandler) Register(router *gin.RouterGroup) {
	router.GET("/", h.ListOrderHistories)
	router.GET("/:id", h.GetOrderHistory)
}

func (h *OrderHistoryHandler) GroupRouterPattern() string {
	return "/api/v1/orders/histories"
}

// GetOrderHistory godoc
//
//	@Summary		Get an order history
//	@Description	Get an order history
//	@Tags			order-histories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Order History ID"
//	@Success		200	{object}	response.OrderResponse
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/orders/histories/{id} [get]
func (h *OrderHistoryHandler) GetOrderHistory(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	orderHistory, err := h.service.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	orderHistoryResponse := response.NewOrderHistoryResponse(orderHistory)
	c.JSON(http.StatusOK, orderHistoryResponse)
}

// ListOrderHistories godoc
//
//	@Summary		List order histories
//	@Description	List order histories
//	@Tags			orderHistories
//	@Accept			json
//	@Produce		json
//	@Param			status		query		string	false	"Status name"
//	@Param			order_id	query		uint64	false	"Order ID"
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Success		200			{object}	response.OrderPaginated
//	@Router			/api/v1/orders/histories [get]
func (h *OrderHistoryHandler) ListOrderHistories(c *gin.Context) {
	orderID := c.DefaultQuery("order_id", "0")
	status := c.DefaultQuery("status", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	orderIDUint64, err := strconv.ParseUint(orderID, 10, 64)
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
	orders, total, err := h.service.List(orderIDUint64, &orderStatus, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	responses := response.NewOrderHistoryPaginated(orders, total, pageInt, limitInt)
	c.JSON(http.StatusOK, responses)
}
