package response

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func NewHealthCheckResponse() *HealthCheckResponse {
	return &HealthCheckResponse{
		Message: "Ok",
	}
}
