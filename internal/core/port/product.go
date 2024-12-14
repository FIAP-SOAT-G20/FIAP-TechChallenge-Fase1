package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type ProductRepository interface {
	Insert(product *domain.Product) error
	GetByID(id uint64) (*domain.Product, error)
	GetAll(name string, categoryID, page, limit uint64) ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint64) error
}

type ProductService interface {
	Create(product *domain.Product) error
	GetByID(id uint64) (*domain.Product, error)
	List(name string, categoryID, page, limit uint64) ([]domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint64) error
}
