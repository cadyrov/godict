package godict

type DictionaryRender struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Dictionary map[string]map[int]string
