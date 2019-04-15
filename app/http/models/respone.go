package models

import (
	"encoding/json"
	"net/http"
)

// Response - структура ответа рест
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// Send - метод отправки ответа
func (this *Response) Send(writer http.ResponseWriter, code int, message string, payload interface{}) {
	this.Code = code
	this.Message = message
	this.Payload = payload
	writer.Header().Set("Content-Type", "application/json; charset=utf8")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(this)
}
