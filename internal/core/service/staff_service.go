package service

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/port"
)

type CustomerService struct {
	customerRepository port.ICustomerRepository
}

func NewCustomerService(customerRepository port.ICustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (cs *CustomerService) Create(customer *domain.Customer) error {
	return cs.customerRepository.Insert(customer)
}

func (cs *CustomerService) GetByID(id uint64) (*domain.Customer, error) {
	return cs.customerRepository.GetByID(id)
}

func (cs *CustomerService) GetByCPF(cpf string) (*domain.Customer, error) {
	return cs.customerRepository.GetByCPF(cpf)
}

func (cs *CustomerService) List(name string, page, limit int) ([]domain.Customer, int64, error) {
	return cs.customerRepository.GetAll(name, page, limit)
}

func (cs *CustomerService) Update(customer *domain.Customer) error {
	_, err := cs.customerRepository.GetByID(customer.ID)
	if err != nil {
		return domain.ErrNotFound
	}

	return cs.customerRepository.Update(customer)
}

func (cs *CustomerService) Delete(id uint64) error {
	_, err := cs.customerRepository.GetByID(id)
	if err != nil {
		return domain.ErrNotFound
	}

	return cs.customerRepository.Delete(id)
}
