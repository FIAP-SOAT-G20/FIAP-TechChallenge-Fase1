package response

type Paginated struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

func NewPagination(total int64, page, limit int) Paginated {
	return Paginated{
		Total: total,
		Page:  page,
		Limit: limit,
	}
}
