package models

type Pagination struct {
	Page      int   `json:"page"`
	Total     int64 `json:"total"`
	Page_Size int   `json:"page_size"`
}

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
