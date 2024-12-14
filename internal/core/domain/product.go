package domain

import "time"

type Product struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Active      bool      `json:"active"`
	CategoryID  uint64    `json:"categoryID"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	// TODO: add staff id
}
