package port

import "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/core/domain"

type IOrderRepository interface {
	Insert(order *domain.Order) error
	GetByID(id uint64) (*domain.Order, error)
	GetAll(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error)
	Update(order *domain.Order) error
	Delete(id uint64) error
}

type IOrderProductRepository interface {
	Insert(orderProduct *domain.OrderProduct) error
	InsertMany(orderProduct []domain.OrderProduct) error
	GetByID(id uint64) (*domain.OrderProduct, error)
	GetAll(orderId, productId uint64, page, limit int) ([]domain.OrderProduct, int64, error)
	Update(order *domain.OrderProduct) error
	Delete(id uint64) error
}

type IOrderHistoryRepository interface {
	Insert(orderHistory *domain.OrderHistory) error
	GetByID(id uint64) (*domain.OrderHistory, error)
	GetAll(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error)
	Delete(id uint64) error
}

type IOrderService interface {
	Create(order *domain.Order) error
	GetByID(id uint64) (*domain.Order, error)
	List(customerID uint64, status *domain.OrderStatus, page, limit int) ([]domain.Order, int64, error)
	Update(order *domain.Order) error
	Delete(id uint64) error
	//UpdateStatus(id uint64, status domain.OrderStatus) error
}

type IOrderHistoryService interface {
	Create(orderID uint64, staffID *uint64, status domain.OrderStatus) error
	GetByID(id uint64) (*domain.OrderHistory, error)
	List(orderID uint64, status *domain.OrderStatus, page, limit int) ([]domain.OrderHistory, int64, error)
	Delete(id uint64) error
}

type IOrderProductService interface {
	Create(orderProduct *domain.OrderProduct) error
	GetByID(id uint64) (*domain.OrderProduct, error)
	List(orderID uint64, page, limit int) ([]domain.OrderProduct, int64, error)
	Update(orderProduct *domain.OrderProduct) error
	Delete(id uint64) error
}
