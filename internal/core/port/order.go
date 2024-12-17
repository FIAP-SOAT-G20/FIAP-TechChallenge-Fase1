package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type OrderRepository interface {
	Insert(order *domain.Order) error
	GetByID(id uint64) (*domain.Order, error)
	GetAll(clientID uint64, page, limit int) ([]domain.Order, int64, error)
	Update(order *domain.Order) error
	Delete(id uint64) error
}

type OrderProductRepository interface {
	Insert(orderProduct *domain.OrderProduct) error
	InsertMany(orderProduct []domain.OrderProduct) error
}

type OrderHistoryRepository interface {
	Insert(orderHistory *domain.OrderHistory) error
}

type OrderService interface {
	Create(order *domain.Order) error
	GetByID(id uint64) (*domain.Order, error)
	List(clientID uint64, page, limit int) ([]domain.Order, int64, error)
	Update(order *domain.Order) error
	Delete(id uint64) error
	UpdateStatus(id uint64, status domain.OrderStatus) error
	AddProduct(id uint64, product *domain.Product) error
	AddProducts(id uint64, product []domain.Product) error
}
