package handler

import (
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

// CategoryHandler represents the HTTP handler for category-related requests
type CategoryHandler struct {
	service port.ICategoryService
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewCategoryHandler(service port.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		service,
	}
}

func (h *CategoryHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateCategory)
	router.GET("/", h.ListCategories)
	router.GET("/:id", h.GetCategory)
	router.PUT("/:id", h.UpdateCategory)
	router.DELETE("/:id", h.DeleteCategory)
}

func (h *CategoryHandler) GroupRouterPattern() string {
	return "/api/v1/categories"
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	create a new category with name
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			createCategoryRequest	body		createCategoryRequest	true	"Create Category Request"
//	@Success		200						{object}	response.CategoryResponse		"Category created"
//	@Failure		400						{object}	response.ErrorResponse	"Validation error"
//	@Failure		401						{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		403						{object}	response.ErrorResponse	"Forbidden error"
//	@Failure		404						{object}	response.ErrorResponse	"Data not found error"
//	@Failure		409						{object}	response.ErrorResponse	"Data conflict error"
//	@Failure		500						{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories [post]
func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	category := req.ToDomain()

	err := ch.service.Create(category)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusCreated, rsp)
}

// GetCategory godoc
//
//	@Summary		Get a category
//	@Description	get a category by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64		true	"Category ID"
//	@Success		200	{object}	response.CategoryResponse	"Category retrieved"
//	@Failure		400	{object}	response.ErrorResponse		"Validation error"
//	@Failure		404	{object}	response.ErrorResponse		"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/api/v1/categories/{id} [get]
func (ch *CategoryHandler) GetCategory(ctx *gin.Context) {
	var req request.GetCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	category, err := ch.service.GetByID(req.ID)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

// ListCategories godoc
//
//	@Summary		List categories
//	@Description	List categories with pagination
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Name"
//	@Param			page	query		int		false	"Page"
//	@Param			limit	query		int		false	"Limit"
//	@Success		200		{object}	response.CategoriesPaginated	"List of categories"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories [get]
func (ch *CategoryHandler) ListCategories(ctx *gin.Context) {
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		response.HandleError(ctx, domain.ErrInvalidQueryParams)
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		response.HandleError(ctx, domain.ErrInvalidQueryParams)
		return
	}

	customers, total, err := ch.service.List(name, pageInt, limitInt)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	customersResponse := response.NewCategoriesPaginated(customers, total, pageInt, limitInt)
	ctx.JSON(http.StatusOK, customersResponse)
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	update a category's name by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint64					true	"Category ID"
//	@Param			updateCategoryRequest	body		request.UpdateCategoryRequest	true	"Update Category Request"
//	@Success		200		{object}	response.CategoryResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories/{id} [put]
//	@Security		BearerAuth
func (ch *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ValidationError(ctx, err)
		return
	}

	category := req.ToDomain(id)

	err = ch.service.Update(category)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

// DeleteCategory godoc
//
//	@Summary		Delete a category
//	@Description	Delete a category by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64			true	"Category ID"
//	@Success		204	{object}	string		"Category deleted"
//	@Failure		400	{object}	response.ErrorResponse	"Validation error"
//	@Failure		401	{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	response.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories/{id} [delete]
func (ch *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	var req request.DeleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	err := ch.service.Delete(req.ID)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
