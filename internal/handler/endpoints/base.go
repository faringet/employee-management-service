package endpoints

type BaseIdRequest struct {
	ID int `json:"id"`
}

type BaseIdResponse struct {
	ID int `json:"id"`
}

type PaginationRequest struct {
	SortBy             string
	SortType           string
	SearchBy           string
	SearchValue        string
	SearchLogicOpeator string
	Limit              int
	Offset             int
}
