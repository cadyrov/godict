package godict

import (
	"encoding/json"
	"github.com/cadyrov/goerr"
	"io"
	"io/ioutil"
	"net/http"
)

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

func ParseBody(r io.ReadCloser, data interface{}) goerr.IError {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return goerr.New(err.Error()).Http(http.StatusBadRequest)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()
	err = json.Unmarshal(body, data)
	if err != nil {
		return goerr.New(err.Error()).Http(http.StatusBadRequest)
	}
	return nil
}

func Ok(message interface{}, data interface{}, pagination interface{}) *Response {
	return &Response{HttpCode: http.StatusOK, Message: message, Data: data, Pagination: pagination}
}

func Error(e goerr.IError) *Response {
	httpCode := http.StatusInternalServerError
	if e.GetCode() >= http.StatusBadRequest && e.GetCode() <= http.StatusNetworkAuthenticationRequired {
		httpCode = e.GetCode()
	}
	return &Response{HttpCode: httpCode, Error: e}
}

func Send(writer http.ResponseWriter, response *Response) {
	SendJson(writer, response.HttpCode, response)
}

func SendError(writer http.ResponseWriter, e goerr.IError) {
	Send(writer, Error(e))
}

func SendOk(writer http.ResponseWriter, message interface{}, data interface{}, pagination interface{}) {
	Send(writer, Ok(message, data, pagination))
}

func SendJson(writer http.ResponseWriter, httpCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpCode)
	if data == nil {
		return
	}
	body, err := json.Marshal(data)
	if err != nil {
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
		return
	}
	_, err = writer.Write(body)
	if err != nil {
		panic(err)
	}
}
