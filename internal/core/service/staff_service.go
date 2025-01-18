package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type StaffService struct {
	staffRepository port.IStaffRepository
}

func NewStaffService(staffRepository port.IStaffRepository) *StaffService {
	return &StaffService{
		staffRepository: staffRepository,
	}
}

func (ss *StaffService) Create(staff *domain.Staff) error {
	return ss.staffRepository.Insert(staff)
}

func (ss *StaffService) GetByID(id uint64) (*domain.Staff, error) {
	return ss.staffRepository.GetByID(id)
}

func (ss *StaffService) List(name string, page, limit int) ([]domain.Staff, int64, error) {
	return ss.staffRepository.GetAll(name, page, limit)
}

func (ss *StaffService) Update(staff *domain.Staff) error {
	_, err := ss.staffRepository.GetByID(staff.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	return ss.staffRepository.Update(staff)
}

func (ss *StaffService) Delete(id uint64) error {
	_, err := ss.staffRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return ss.staffRepository.Delete(id)
}
