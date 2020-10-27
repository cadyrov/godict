package godict

const (
	ENi = 45
	RUi = 643

	ENLocale = Locale(ENi)
	RULocale = Locale(RUi)
)

type Locale int

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

func (d *Dictionary) Render() map[string]*[]DictionaryRender {
	result := make(map[string]*[]DictionaryRender)

	for key, value := range *d {
		dr := make([]DictionaryRender, 0, len(value))

		for rKey, rVal := range value {
			dr = append(dr, DictionaryRender{
				ID:   rKey,
				Name: rVal,
			})
		}

		result[key] = &dr
	}

	return result
}
