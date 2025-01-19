package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

// CreateStaffRequest contains the request to create a staff
type CreateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"COOK"`
}

// UpdateStaffRequest contains the request to update a staff
type UpdateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"ATTENDANT"`
}
