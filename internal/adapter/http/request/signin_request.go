package request

// SignInRequest represents the request body for logging in a user
type SignInRequest struct {
	Cpf string `json:"cpf" binding:"required" example:"000.000.000-00"`
}
