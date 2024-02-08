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
	ActivePage   int `json:"active_page"`
	TotalCount   int `json:"total_count"`
	CountPerPage int `json:"count_per_page"`
	TotalPages   int `json:"total_pages"`
}

type CountResponse struct {
	Count int64  `json:"count"`
	Test  string `json:"test"`
}