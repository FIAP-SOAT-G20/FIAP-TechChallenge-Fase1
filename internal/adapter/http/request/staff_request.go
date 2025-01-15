package request

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type CreateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"COOK, ATTENDANT or MANAGER"`
}

type UpdateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"COOK, ATTENDANT or MANAGER"`
}
