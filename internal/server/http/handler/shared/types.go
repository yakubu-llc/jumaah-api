package shared

type PaginationRequest struct {
	Cursor int `query:"cursor" required:"false" default:"0"`
	Limit  int `query:"limit" required:"false" default:"10"`
}

type PathIDParam struct {
	ID int `path:"id"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type PaginationResponse struct {
	Cursor  *int `json:"cursor"`
	HasMore bool `json:"hasMore"`
}
