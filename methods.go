package godict

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/cadyrov/goerr/v2"
	"gopkg.in/yaml.v2"
)

var (
	ErrDictionaryNotFound = errors.New("dictionary_not_found")
)

func (d *Dictionary) DictionaryRender(dictionaryName string,
	dictionaryID int) (res DictionaryRender, e goerr.IError) {
	if !d.IsKeyExists(dictionaryName) {
		e = goerr.NotFound(ErrDictionaryNotFound)

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
		e = goerr.NotFound(ErrDictionaryNotFound)

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
		return goerr.BadRequest(err)
	}

	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = json.Unmarshal(body, data)
	if err != nil {
		return goerr.BadRequest(err)
	}

	return nil
}

func ParseBodyReturned(r io.ReadCloser, data interface{})(body []byte ,e  goerr.IError) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, goerr.BadRequest(err)
	}

	defer func() {
		err := r.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = json.Unmarshal(body, data)
	if err != nil {
		return body, goerr.BadRequest(err)
	}

	return body, nil
}

func Ok(message interface{}, data interface{}, pagination interface{}) *Response {
	return &Response{HTTPCode: http.StatusOK, Message: message, Data: data, Pagination: pagination}
}

func Error(e goerr.IError) *Response {
	httpCode := http.StatusInternalServerError

	if e.Code() >= http.StatusBadRequest && e.Code() <= http.StatusNetworkAuthenticationRequired {
		httpCode = e.Code()
	}

	return &Response{HTTPCode: httpCode, Error: e.Error()}
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
		e = goerr.Internal(err)
	}

	return data, e
}

func YamlMarshal(val interface{}) (data []byte, e goerr.IError) {
	data, err := yaml.Marshal(val)
	if err != nil {
		e = goerr.Internal(err)
	}

	return data, e
}

func Unmarshal(data []byte, val interface{}) (e goerr.IError) {
	if data == nil {
		return
	}

	err := json.Unmarshal(data, val)
	if err != nil {
		e = goerr.Internal(err)
	}

	return e
}

func YamlUnmarshal(data []byte, val interface{}) (e goerr.IError) {
	if data == nil {
		return
	}

	err := yaml.Unmarshal(data, val)
	if err != nil {
		e = goerr.Internal(err)
	}

	return e
}
