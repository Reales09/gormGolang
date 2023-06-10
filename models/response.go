package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json: "status"`
	Data          interface{} `json: "data"`
	Message       string      `json: "message"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: rw,
		contentType:   "Aplication/json",
	}
}

func (resp *Response) Send() {
	resp.responseWrite.Header().Set("contentType", resp.contentType)
	resp.responseWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.responseWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {

	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource No found"
}

func SendFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = " UnprocessableEntity No found"
}

func SendUnProcessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
