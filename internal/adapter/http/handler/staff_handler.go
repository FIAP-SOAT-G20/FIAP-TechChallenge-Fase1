package handler

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/request"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/response"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type StaffHandler struct {
	staffService port.IStaffService
}

func NewStaffHandler(staffService port.IStaffService) *StaffHandler {
	return &StaffHandler{staffService: staffService}
}

func (h *StaffHandler) Register(router *gin.RouterGroup) {
	router.POST("/", h.CreateStaff)
	router.GET("/", h.ListStaffs)
	router.GET("/:id", h.GetStaff)
	router.PUT("/:id", h.UpdateStaff)
	router.DELETE("/:id", h.DeleteStaff)
}

func (h *StaffHandler) GroupRouterPattern() string {
	return "/api/v1/staffs"
}

// CreateStaff godoc
//
//	@Summary		Create a staff
//	@Description	Create a staff
//	@Tags			staffs, sign-up
//	@Accept			json
//	@Produce		json
//	@Param			staff	body		request.CreateStaffRequest	true	"Staff"
//	@Success		201			{object}	response.StaffResponse
//	@Failure		400			{object}	response.ErrorResponse	"Validation error"
//	@Failure		500			{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/staffs [post]
func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var req request.CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	staff := &domain.Staff{
		Name: req.Name,
		Role: req.Role,
	}

	err := h.staffService.Create(staff)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	staffResponse := response.NewStaffResponse(staff)
	c.JSON(http.StatusCreated, staffResponse)
}

// ListStaffs godoc
//
//	@Summary		List staffs
//	@Description	List staffs
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	false	"Name"
//	@Param			page	query		int		false	"Page"
//	@Param			limit	query		int		false	"Limit"
//	@Success		200		{object}	response.StaffsPaginated
//	@Failure		500		{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/staffs [get]
func (h *StaffHandler) ListStaffs(c *gin.Context) {
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

	staffs, total, err := h.staffService.List(name, pageInt, limitInt)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	staffsResponse := response.NewStaffsPaginated(staffs, total, pageInt, limitInt)
	c.JSON(http.StatusOK, staffsResponse)
}

// GetStaff godoc
//
//	@Summary		Get a staff
//	@Description	Get a staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Staff ID"
//	@Success		200	{object}	response.StaffResponse
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/staffs/{id} [get]
func (h *StaffHandler) GetStaff(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	staff, err := h.staffService.GetByID(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	staffResponse := response.NewStaffResponse(staff)
	c.JSON(http.StatusOK, staffResponse)
}

// UpdateStaff godoc
//
//	@Summary		Update a staff
//	@Description	Update a staff
//	@Tags			customers
//	@Accept			json
//	@Produce		json
//	@Param			id			path		uint64						true	"Staff ID"
//	@Param			customer	body		request.UpdateStaffRequest	true	"Staff"
//	@Success		200			{object}	response.CustomerResponse
//	@Failure		400			{object}	response.ErrorResponse	"Validation error"
//	@Failure		404			{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500			{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/staffs/{id} [put]
func (h *StaffHandler) UpdateStaff(c *gin.Context) {
	id := c.Param("id")

	var req request.UpdateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	staff := &domain.Staff{
		ID:        idUint64,
		Name:      req.Name,
		Role:      req.Role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = h.staffService.Update(staff)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	staffResponse := response.NewStaffResponse(staff)
	c.JSON(http.StatusOK, staffResponse)
}

// DeleteStaff godoc
//
//	@Summary		Delete a staff
//	@Description	Delete a staff
//	@Tags			staffs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	uint64	true	"Staff ID"
//	@Success		204
//	@Failure		404	{object}	response.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	response.ErrorResponse	"Internal server error"
//	@Router			/api/v1/staffs/{id} [delete]
func (h *StaffHandler) DeleteStaff(c *gin.Context) {
	id := c.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		response.HandleError(c, domain.ErrInvalidParam)
		return
	}

	err = h.staffService.Delete(idUint64)
	if err != nil {
		response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
