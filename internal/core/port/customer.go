package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type ICustomerRepository interface {
	Insert(customer *domain.Customer) error
	GetByID(id uint64) (*domain.Customer, error)
	GetByCPF(cpf string) (*domain.Customer, error)
	GetAll(name string, page, limit int) ([]domain.Customer, int64, error)
	Update(customer *domain.Customer) error
	Delete(id uint64) error
}

type ICustomerService interface {
	Create(customer *domain.Customer) error
	GetByID(id uint64) (*domain.Customer, error)
	List(name string, page, limit int) ([]domain.Customer, int64, error)
	Update(customer *domain.Customer) error
	Delete(id uint64) error
}
