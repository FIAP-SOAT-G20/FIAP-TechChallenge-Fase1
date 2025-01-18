package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db}
}

func (r *StaffRepository) Insert(customer *domain.Staff) error {
	return r.db.Create(customer).Error
}

func (r *StaffRepository) GetByID(id uint64) (*domain.Staff, error) {
	var customer domain.Staff

	err := r.db.First(&customer, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	return &customer, err
}

func (r *StaffRepository) GetAll(name string, page, limit int) ([]domain.Staff, int64, error) {
	var customers []domain.Staff
	var count int64

	query := r.db.Model(&domain.Staff{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Count(&count)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&customers).Error

	return customers, count, err
}

func (r *StaffRepository) Update(customer *domain.Staff) error {
	return r.db.Save(customer).Error
}

func (r *StaffRepository) Delete(id uint64) error {
	return r.db.Delete(&domain.Staff{}, id).Error
}
