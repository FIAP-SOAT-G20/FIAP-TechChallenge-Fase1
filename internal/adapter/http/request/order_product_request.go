package request

type OrderProductRequest struct {
	Quantity uint32 `json:"quantity" binding:"required" example:"10"`
}

type OrderProductPathParam struct {
	OrderID   uint64 `uri:"orderID" binding:"required"`
	ProductID uint64 `uri:"productID" binding:"required"`
}
