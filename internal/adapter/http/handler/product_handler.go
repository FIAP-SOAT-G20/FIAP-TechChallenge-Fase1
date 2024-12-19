package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type ProductHandler struct {
	service port.IProductService
}

func NewProductHandler(service port.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateProduct)
	router.GET("/", h.ListProducts)
	router.GET("/:id", h.GetProduct)
	router.PUT("/:id", h.UpdateProduct)
	router.DELETE("/:id", h.DeleteProduct)
}

func (h *ProductHandler) GroupRouterPattern() string {
	return "/api/v1/products"
}

type createProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"BK Mega Stacker 2.0"`
	Description string  `json:"description" binding:"required" example:"The best burger in the world"`
	Price       float64 `json:"price" binding:"required" example:"29.90"`
	CategoryID  uint64  `json:"categoryID" binding:"required" example:"1"`
}

// CreateProduct godoc
//
//	@Summary		Create a product
//	@Description	Create a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		createProductRequest	true	"ProductResponse"
//	@Success		201		{object}	response.ProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if err := h.service.Create(product); err != nil {
		response.HandleError(c, err)
		return
	}

	productResponse := response.NewProductResponse(product)
	c.JSON(http.StatusCreated, productResponse)
}

// GetProduct godoc
//
//	@Summary		Get a product
//	@Description	Get a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ProductResponse ID"
//	@Success		200	{object}	response.ProductResponse
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	product, err := h.service.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	productResponse := response.NewProductResponse(product)
	c.JSON(http.StatusOK, productResponse)
}

// ListProducts godoc
//
//	@Summary		List products
//	@Description	List products
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			name		query		string	false	"ProductResponse name"
//	@Param			categoryID	query		uint64	false	"Category ID"
//	@Param			page		query		int		false	"Page"
//	@Param			limit		query		int		false	"Limit"
//	@Success		200			{object}	response.ProductPaginated
//	@Router			/api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	name := c.Query("name")
	categoryID := c.DefaultQuery("categoryID", "0")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	categoryIDUint64, err := strconv.ParseUint(categoryID, 10, 64)
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

	products, total, err := h.service.List(name, categoryIDUint64, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	productResponses := response.NewProductPaginated(products, total, pageInt, limitInt)
	c.JSON(http.StatusOK, productResponses)
}

type UpdateProduct struct {
	Name        string  `json:"name" example:"McDonald's Big Mac"`
	Description string  `json:"description" example:"The best burger in the world"`
	Price       float64 `json:"price" example:"29.90"`
	CategoryID  uint64  `json:"categoryID" example:"1"`
}

// UpdateProduct godoc
//
//	@Summary		Update a product
//	@Description	Update a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"ProductResponse ID"
//	@Param			product	body		UpdateProduct	true	"ProductResponse"
//	@Success		200		{object}	response.ProductResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	var req UpdateProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	product := &domain.Product{
		ID:          idUint64,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if err := h.service.Update(product); err != nil {
		response.HandleError(c, err)
		return
	}

	productResponse := response.NewProductResponse(product)
	c.JSON(http.StatusOK, productResponse)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product
//	@Description	Delete a product
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ProductResponse ID"
//	@Success		204	{object}	string
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
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
