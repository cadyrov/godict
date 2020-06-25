package godict

import "github.com/cadyrov/goerr"

func (d *Dictionary) DictionaryRender(dictionaryName string, dictionaryId int) (res DictionaryRender, e goerr.IError) {
	if !d.IsKeyExists(dictionaryName) {
		e = goerr.New("dictionary_not_found")
		return
	}
	res.Id = dictionaryId
	res.Name = (*d)[dictionaryName][dictionaryId]
	return
}

func (d *Dictionary) IsKeyExists(key string) (res bool) {
	_, res = (*d)[key]
	return
}

func (d *Dictionary) DictionaryIdsInterface(dictionaryName string) (res []interface{}, e goerr.IError) {
	if !d.IsKeyExists(dictionaryName) {
		e = goerr.New("dictionary not found")
		return
	}
	for i := range (*d)[dictionaryName] {
		res = append(res, i)
	}
	return
}
