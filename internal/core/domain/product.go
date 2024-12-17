package domain

import "time"

type Product struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	Active      bool
	CategoryID  uint64
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
