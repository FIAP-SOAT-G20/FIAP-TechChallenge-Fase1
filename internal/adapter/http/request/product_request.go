package request

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"BK Mega Stacker 2.0"`
	Description string  `json:"description" binding:"required" example:"The best burger in the world"`
	Price       float32 `json:"price" binding:"required" example:"29.90"`
	CategoryID  uint64  `json:"category_id" binding:"required" example:"1"`
	Active      bool    `json:"active" example:"true" default:"true"`
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name        string  `json:"name" example:"McDonald's Big Mac"`
	Description string  `json:"description" example:"The best burger in the world"`
	Price       float32 `json:"price" example:"29.90"`
	CategoryID  uint64  `json:"category_id" example:"1"`
	Active      bool    `json:"active" example:"true"`
}
