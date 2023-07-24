package models

type Pagination struct {
	Page        int   `json:"page"`
	Page_Size   int   `json:"page_size"`
	Total_Pages int64 `json:"total_pages"`
}

type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Token      string      `json:"token,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
