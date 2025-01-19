package request

// OrderProductRequest represents a request body for ordering a product
type OrderProductRequest struct {
	Quantity uint32 `json:"quantity" binding:"required" example:"10"`
}

// OrderProductPathParam represents a request path parameter for ordering a product
type OrderProductPathParam struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}
