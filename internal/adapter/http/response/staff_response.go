package response

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type StaffResponse struct {
	ID   uint64      `json:"id"`
	Name string      `json:"name"`
	Role domain.Role `json:"role"`
}

func NewStaffResponse(staff *domain.Staff) *StaffResponse {
	if staff == nil {
		return nil
	}

	return &StaffResponse{
		ID:   staff.ID,
		Name: staff.Name,
		Role: staff.Role,
	}
}

type StaffsPaginated struct {
	Paginated
	Staffs []StaffResponse `json:"staffs"`
}

func NewStaffsPaginated(staffs []domain.Staff, total int64, page int, limit int) *StaffsPaginated {
	staffResponses := make([]StaffResponse, 0, len(staffs))
	for _, staff := range staffs {
		staffResponse := NewStaffResponse(&staff)
		if staffResponse != nil {
			staffResponses = append(staffResponses, *staffResponse)
		}
	}

	return &StaffsPaginated{
		Paginated: Paginated{
			Total: total,
			Page:  page,
			Limit: limit,
		},
		Staffs: staffResponses,
	}
}
