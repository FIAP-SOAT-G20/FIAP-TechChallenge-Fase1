package domain

type Product struct {
	ID          uint64
	Name        string
	Description string
	Price       float64
	Active      bool
	CategoryID  uint64
	Category    Category
	CreatedAt   string
	UpdateAt    string
	// TODO: add staff id
}
