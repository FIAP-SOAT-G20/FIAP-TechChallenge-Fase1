package repository

import (
	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) Insert(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) GetByID(id uint64) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

func (r *CustomerRepository) GetByCPF(cpf string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.Where("cpf = ?", cpf).First(&customer).Error
	return &customer, err
}

func (r *CustomerRepository) GetAll(name string, page, limit int) ([]domain.Customer, int64, error) {
	var customers []domain.Customer
	var count int64

	query := r.db.Model(&domain.Customer{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Count(&count)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&customers).Error

	return customers, count, err
}

func (r *CustomerRepository) Update(customer *domain.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Customer{}, id).Error
}
