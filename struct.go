package godict

type DictionaryRender struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Dictionary map[string]map[int]string

// Pagination
type Pagination struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

type Response struct {
	HttpCode   int         `json:"-"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
