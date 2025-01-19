package domain

import (
	"strings"
	"time"
)

type OrderStatus string

const (
	UNDEFINDED OrderStatus = "UNDEFINDED"
	OPEN       OrderStatus = "OPEN"
	CANCELLED  OrderStatus = "CANCELLED"
	PENDING    OrderStatus = "PENDING"
	RECEIVED   OrderStatus = "RECEIVED"
	PREPARING  OrderStatus = "PREPARING"
	READY      OrderStatus = "READY"
	COMPLETED  OrderStatus = "COMPLETED"
)

func (o OrderStatus) ToString() string {
	return string(o)
}

func (o OrderStatus) From(status string) OrderStatus {
	switch strings.ToUpper(status) {
	case "OPEN":
		return OPEN
	case "CANCELLED":
		return CANCELLED
	case "PENDING":
		return PENDING
	case "RECEIVED":
		return RECEIVED
	case "PREPARING":
		return PREPARING
	case "READY":
		return READY
	case "COMPLETED":
		return COMPLETED
	default:
		return UNDEFINDED
	}
}

var OrderStatusTransitions = map[OrderStatus][]OrderStatus{
	OPEN:      {CANCELLED, PENDING},
	CANCELLED: {},
	PENDING:   {OPEN, RECEIVED},
	RECEIVED:  {PREPARING},
	PREPARING: {READY},
	READY:     {COMPLETED},
	COMPLETED: {},
}

func CanTransitionTo(oldStatus, newStatus OrderStatus) bool {
	allowedStatuses := OrderStatusTransitions[oldStatus]
	for _, allowedStatus := range allowedStatuses {
		if newStatus == allowedStatus {
			return true
		}
	}
	return false
}

func StatusTransitionNeedsStaffID(newStatus OrderStatus) bool {
	switch newStatus {
	case OPEN:
		return false
	case CANCELLED:
		return false
	case PENDING:
		return false
	case RECEIVED:
		return false
	case PREPARING:
		return true
	case READY:
		return true
	case COMPLETED:
		return true
	default:
		return false
	}
}

type Order struct {
	ID            uint64
	CustomerID    uint64
	TotalBill     float32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        OrderStatus
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
	StaffID   *uint64
	Status    OrderStatus
	CreatedAt time.Time
	Order     Order
	Staff     *Staff
}
