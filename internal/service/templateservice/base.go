package templateservice

type BaseIdResponse struct {
	ID int
}

type BaseIdRequest struct {
	ID int
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
	ActivePage   int `json:"active_page"`
	TotalCount   int `json:"total_count"`
	CountPerPage int `json:"count_per_page"`
	TotalPages   int `json:"total_pages"`
}
