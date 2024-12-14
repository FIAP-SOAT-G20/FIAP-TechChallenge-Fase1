package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type ProductHandler struct {
	service port.ProductService
}

func NewProductHandler(service port.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateProduct)
}

type createProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"BK Mega Stacker 2.0"`
	Description string  `json:"description" binding:"required" example:"The best burger in the world"`
	Price       float64 `json:"price" binding:"required" example:"29.90"`
	CategoryID  uint64  `json:"categoryID" binding:"required" example:"1"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := &domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	if err := h.service.Create(product); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
