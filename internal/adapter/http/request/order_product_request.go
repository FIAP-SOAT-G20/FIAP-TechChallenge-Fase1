package request

type CreateOrderProductRequest struct {
	OrderID   uint64 `json:"order_id" binding:"required" example:"1"`
	ProductID uint64 `json:"product_id" binding:"required" example:"1"`
	Quantity  uint32 `json:"quantity" binding:"required" example:"10"`
}

type UpdateOrderProductRequest struct {
	OrderID   uint64 `json:"order_id" binding:"required" example:"1"`
	ProductID uint64 `json:"product_id" binding:"required" example:"1"`
	Quantity  uint32 `json:"quantity" binding:"required" example:"10"`
}

type DeleteOrderProductRequest struct {
	OrderID   uint64 `json:"order_id" binding:"required" example:"1"`
	ProductID uint64 `json:"product_id" binding:"required" example:"1"`
}
