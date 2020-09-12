package godict

type DictionaryRender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Dictionary map[string]map[int]string

type Pagination struct {
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	Total int `json:"total,omitempty"`
}

type Response struct {
	HTTPCode   int         `json:"-"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
