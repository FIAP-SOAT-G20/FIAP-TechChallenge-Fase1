package request

type CreateOrderRequest struct {
	CustomerID uint64 `json:"customer_id" binding:"required" example:"1"`
}
