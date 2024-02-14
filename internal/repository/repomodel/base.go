package repomodel

type BaseIdRequest struct {
	ID int32 `json:"id"`
}

type BaseIdResponse struct {
	ID int `json:"id"`
}

type PaginationRequest struct {
	SortBy             string
	SortType           string
	SearchBy           []string
	SearchValue        []string
	SearchLogicOpeator string
	Limit              int
	Offset             int
}

type PaginationResponse struct {
	ActivePage   int `sql:"active_page"`
	TotalCount   int `sql:"total_count"`
	CountPerPage int `sql:"count_per_page"`
	TotalPages   int `sql:"total_pages"`
}

type CountResponse struct {
	Count int64 `sql:"count"`
}
