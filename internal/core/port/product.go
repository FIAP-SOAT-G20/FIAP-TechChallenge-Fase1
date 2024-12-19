package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type IProductRepository interface {
	Insert(product *domain.Product) error
	GetByID(id uint64) (*domain.Product, error)
	GetAll(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error)
	Update(product *domain.Product) error
	Delete(id uint64) error
}

type IProductService interface {
	Create(product *domain.Product) error
	GetByID(id uint64) (*domain.Product, error)
	List(name string, categoryID uint64, page, limit int) ([]domain.Product, int64, error)
	Update(product *domain.Product) error
	Delete(id uint64) error
}
