package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type CustomerHandler struct {
	customerService port.CustomerService
}

func NewCustomerHandler(customerService port.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateCustomer)
	router.GET("/", h.ListCustomers)
	router.GET("/:id", h.GetCustomer)
	router.PUT("/:id", h.UpdateCustomer)
	router.DELETE("/:id", h.DeleteCustomer)
}

type createCustomerRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"johndoe@contact.com"`
	CPF   string `json:"cpf" binding:"required" example:"123.456.789-00"`
}

// CreateCustomer godoc
//
//	@Summary		Create a customer
//	@Description	Create a customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			customer	body		createCustomerRequest	true	"Customer"
//	@Success		201			{object}	response.CustomerResponse
//	@Failure		400			{object}	response.ErrorResponse	"Validation error"
//	@Failure		500			{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req createCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	customer := &domain.Customer{
		Name:  req.Name,
		Email: req.Email,
		CPF:   req.CPF,
	}

	err := h.customerService.Create(customer)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	customerResponse := response.NewCustomerResponse(customer)
	c.JSON(http.StatusCreated, customerResponse)
}

// ListCustomers godoc
//
//	@Summary		List customers
//	@Description	List customers
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Name"
//	@Param			page	query		int		false	"Page"
//	@Param			limit	query		int		false	"Limit"
//	@Success		200		{object}	response.CustomersPaginated
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	name := c.Query("name")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

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

	customers, total, err := h.customerService.List(name, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	customersResponse := response.NewCustomersPaginated(customers, total, pageInt, limitInt)
	c.JSON(http.StatusOK, customersResponse)
}

// GetCustomer godoc
//
//	@Summary		Get a customer
//	@Description	Get a customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Customer ID"
//	@Success		200	{object}	response.CustomerResponse
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	customer, err := h.customerService.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	customerResponse := response.NewCustomerResponse(customer)
	c.JSON(http.StatusOK, customerResponse)
}

type updateCustomerRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"johndoe@email.com"`
	CPF   string `json:"cpf" example:"123.456.789-00"`
}

// UpdateCustomer godoc
//
//	@Summary		Update a customer
//	@Description	Update a customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id			path		uint64					true	"Customer ID"
//	@Param			customer	body		updateCustomerRequest	true	"Customer"
//	@Success		200			{object}	response.CustomerResponse
//	@Failure		400			{object}	response.ErrorResponse	"Validation error"
//	@Failure		404			{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500			{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	var req updateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	customer := &domain.Customer{
		ID:    idUint64,
		Name:  req.Name,
		Email: req.Email,
		CPF:   req.CPF,
	}

	err = h.customerService.Update(customer)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	customerResponse := response.NewCustomerResponse(customer)
	c.JSON(http.StatusOK, customerResponse)
}

// DeleteCustomer godoc
//
//	@Summary		Delete a customer
//	@Description	Delete a customer
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id	path	uint64	true	"Customer ID"
//	@Success		204
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	err = h.customerService.Delete(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
