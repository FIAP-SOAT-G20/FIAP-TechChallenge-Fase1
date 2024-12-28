package domain

import "time"

type OrderStatus string

const (
	OPEN      OrderStatus = "OPEN"
	CANCELLED OrderStatus = "CANCELLED"
	PENDING   OrderStatus = "PENDING"
	RECEIVED  OrderStatus = "RECEIVED"
	PREPARING OrderStatus = "PREPARING"
	READY     OrderStatus = "READY"
	COMPLETED OrderStatus = "COMPLETED"
)

type Order struct {
	ID            uint64
	CustomerID    uint64
	TotalBill     float32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Payment       Payment
	Customer      Customer
	OrderProducts []OrderProduct
}

type OrderProduct struct {
	OrderID   uint64
	ProductID uint64
	Price     float32
	Quantity  uint32
	CreatedAt time.Time
	UpdatedAt time.Time
	Order     Order
	Product   Product
}

type OrderHistory struct {
	ID        uint64
	OrderID   uint64
	StaffID   uint64
	Status    OrderStatus
	CreatedAt time.Time
	Order     Order
	Staff     Staff
}
