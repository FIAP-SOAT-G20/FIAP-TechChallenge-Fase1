package handler

import (
	"net/http"
	"strconv"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
	"github.com/gin-gonic/gin"
)

// CategoryHandler represents the HTTP handler for category-related requests
type CategoryHandler struct {
	svc port.ICategoryService
}

// NewCategoryHandler creates a new CategoryHandler instance
func NewCategoryHandler(svc port.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		svc,
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

// createCategoryRequest represents a request body for creating a new category
type createCategoryRequest struct {
	Name string `json:"name" binding:"required" example:"Foods"`
}

// CreateCategory godoc
//
//	@Summary		Create a new category
//	@Description	create a new category with name
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			createCategoryRequest	body		createCategoryRequest	true	"Create category request"
//	@Success		200						{object}	response.CategoryResponse		"Category created"
//	@Failure		400						{object}	response.ErrorResponse	"Validation error"
//	@Failure		401						{object}	response.ErrorResponse	"Unauthorized error"
//	@Failure		403						{object}	response.ErrorResponse	"Forbidden error"
//	@Failure		404						{object}	response.ErrorResponse	"Data not found error"
//	@Failure		409						{object}	response.ErrorResponse	"Data conflict error"
//	@Failure		500						{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories [post]
func (ch *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	category := &domain.Category{
		Name: req.Name,
	}

	err := ch.svc.Create(category)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusCreated, rsp)
}

// getCategoryRequest represents a request body for retrieving a category
type getCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
}

// GetCategory godoc
//
//	@Summary		Get a category
//	@Description	get a category by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64				true	"Category ID"
//	@Success		200	{object}	response.CategoryResponse	"Category retrieved"
//	@Failure		400	{object}	response.ErrorResponse		"Validation error"
//	@Failure		404	{object}	response.ErrorResponse		"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse		"Internal server error"
//	@Router			/categories/{id} [get]
func (ch *CategoryHandler) GetCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	category, err := ch.svc.GetByID(req.ID)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

// listCategoriesRequest represents a request body for listing categories
type listCategoriesRequest struct {
	Skip  uint64 `form:"skip" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"5"`
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

	customers, total, err := ch.svc.List(name, pageInt, limitInt)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	customersResponse := response.NewCategoriesPaginated(customers, total, pageInt, limitInt)
	ctx.JSON(http.StatusOK, customersResponse)
}

// updateCategoryRequest represents a request body for updating a category
type updateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,required" example:"Beverages"`
}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	update a category's name by id
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id						path		uint64					true	"Category ID"
//	@Param			updateCategoryRequest	body		updateCategoryRequest	true	"Update category request"
//	@Success		200		{object}	response.CategoryResponse
//	@Failure		400		{object}	response.ErrorResponse	"Validation error"
//	@Failure		404		{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/categories/{id} [put]
//	@Security		BearerAuth
func (ch *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var req updateCategoryRequest
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

	category := &domain.Category{
		ID:   id,
		Name: req.Name,
	}

	err = ch.svc.Update(category)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	rsp := response.NewCategoryResponse(category)
	ctx.JSON(http.StatusOK, rsp)
}

// deleteCategoryRequest represents a request body for deleting a category
type deleteCategoryRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1" example:"1"`
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
	var req deleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ValidationError(ctx, err)
		return
	}

	err := ch.svc.Delete(req.ID)
	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
