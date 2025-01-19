package request

// CreateCustomerRequest represents the request to create a customer
type CreateCustomerRequest struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"johndoe@contact.com"`
	CPF   string `json:"cpf" binding:"required" example:"123.456.789-00"`
}

type UpdateCustomerRequest struct {
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"johndoe@email.com"`
	CPF   string `json:"cpf" example:"123.456.789-00"`
}
