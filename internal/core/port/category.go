package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type CategoryRepository interface {
	Insert(category *domain.Category) error
	GetByID(id uint64) (*domain.Category, error)
	GetAll(name string, page, limit uint64) ([]domain.Category, error)
}
