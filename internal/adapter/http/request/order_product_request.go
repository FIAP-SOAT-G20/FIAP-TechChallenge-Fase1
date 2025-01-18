package request

type OrderProductRequest struct {
	Quantity uint32 `json:"quantity" binding:"required" example:"10"`
}

type OrderProductPathParam struct {
	OrderID   uint64 `uri:"order_id" binding:"required"`
	ProductID uint64 `uri:"product_id" binding:"required"`
}
