package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type IStaffRepository interface {
	Insert(staff *domain.Staff) error
	GetByID(id uint64) (*domain.Staff, error)
	GetAll(name string, page, limit int) ([]domain.Staff, int64, error)
	Update(staff *domain.Staff) error
	Delete(id uint64) error
}

type IStaffService interface {
	Create(staff *domain.Staff) error
	GetByID(id uint64) (*domain.Staff, error)
	List(name string, page, limit int) ([]domain.Staff, int64, error)
	Update(staff *domain.Staff) error
	Delete(id uint64) error
}
