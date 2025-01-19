package request

import (
	"time"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

// CreateStaffRequest contains the request to create a staff
type CreateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"COOK"`
}

// ToDomain converts CreateStaffRequest to domain.Staff
func (r CreateStaffRequest) ToDomain() *domain.Staff {
	return &domain.Staff{
		Name: r.Name,
		Role: r.Role,
	}
}

// UpdateStaffRequest contains the request to update a staff
type UpdateStaffRequest struct {
	Name string      `json:"name" binding:"required" example:"John Doe"`
	Role domain.Role `json:"role" binding:"required" example:"ATTENDANT"`
}

// ToDomain converts UpdateStaffRequest to domain.Staff
func (r UpdateStaffRequest) ToDomain(id uint64) *domain.Staff {
	return &domain.Staff{
		ID:        id,
		Name:      r.Name,
		Role:      r.Role,
		UpdatedAt: time.Now(),
	}
}
