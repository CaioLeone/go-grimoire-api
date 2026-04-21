package api

type PaginatedResponse struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}
