package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type ICategoryRepository interface {
	Insert(category *domain.Category) error
	GetByID(id uint64) (*domain.Category, error)
	GetAll(name string, page, limit int) ([]domain.Category, int64, error)
	Update(category *domain.Category) error
	Delete(id uint64) error
}

type ICategoryService interface {
	Create(category *domain.Category) error
	GetByID(id uint64) (*domain.Category, error)
	List(name string, page, limit int) ([]domain.Category, int64, error)
	Update(category *domain.Category) error
	Delete(id uint64) error
}
