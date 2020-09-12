package godict

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cadyrov/goerr"
	"gopkg.in/yaml.v2"
)

func (d *Dictionary) DictionaryRender(dictionaryName string,
	dictionaryID int) (res DictionaryRender, e goerr.IError) {
	if !d.IsKeyExists(dictionaryName) {
		e = goerr.New("dictionary_not_found")

		return
	}

	res.ID = dictionaryID
	res.Name = (*d)[dictionaryName][dictionaryID]

	return res, e
}

func (d *Dictionary) IsKeyExists(key string) (res bool) {
	_, res = (*d)[key]

	return res
}

func (d *Dictionary) DictionaryIdsInterface(dictionaryName string) (res []interface{}, e goerr.IError) {
	if !d.IsKeyExists(dictionaryName) {
		e = goerr.New("dictionary not found")

		return
	}

	for i := range (*d)[dictionaryName] {
		res = append(res, i)
	}

	return res, nil
}

func ParseBody(r io.ReadCloser, data interface{}) goerr.IError {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return goerr.New(err.Error()).HTTP(http.StatusBadRequest)
	}

	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = json.Unmarshal(body, data)
	if err != nil {
		return goerr.New(err.Error()).HTTP(http.StatusBadRequest)
	}

	return nil
}

func Ok(message interface{}, data interface{}, pagination interface{}) *Response {
	return &Response{HTTPCode: http.StatusOK, Message: message, Data: data, Pagination: pagination}
}

func Error(e goerr.IError) *Response {
	httpCode := http.StatusInternalServerError

	if e.GetCode() >= http.StatusBadRequest && e.GetCode() <= http.StatusNetworkAuthenticationRequired {
		httpCode = e.GetCode()
	}

	return &Response{HTTPCode: httpCode, Error: e}
}

func Send(writer http.ResponseWriter, response *Response) {
	SendJSON(writer, response.HTTPCode, response)
}

func SendError(writer http.ResponseWriter, e goerr.IError) {
	Send(writer, Error(e))
}

func SendOk(writer http.ResponseWriter, message interface{}, data interface{},
	pagination interface{}) {
	Send(writer, Ok(message, data, pagination))
}

func SendJSON(writer http.ResponseWriter, httpCode int, data interface{}) {
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

func Marshal(val interface{}) (data []byte, e goerr.IError) {
	data, err := json.Marshal(val)
	if err != nil {
		e = goerr.New(err.Error())
	}

	return data, e
}

func YamlMarshal(val interface{}) (data []byte, e goerr.IError) {
	data, err := yaml.Marshal(val)
	if err != nil {
		e = goerr.New(err.Error())
	}

	return data, e
}

func Unmarshal(data []byte, val interface{}) (e goerr.IError) {
	if data == nil {
		return
	}

	err := json.Unmarshal(data, val)
	if err != nil {
		e = goerr.New(err.Error())
	}

	return e
}

func YamlUnmarshal(data []byte, val interface{}) (e goerr.IError) {
	if data == nil {
		return
	}

	err := yaml.Unmarshal(data, val)
	if err != nil {
		e = goerr.New(err.Error())
	}

	return e
}
