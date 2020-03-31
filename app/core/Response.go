package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SuccessResponse struct {
	Data interface{}
}

type ErrorResponse struct {
	Error string
}

func MakeSuccessResponse(writer http.ResponseWriter, response *SuccessResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	makeResponse(writer, jsonResponse)
}

func MakeErrorResponse(writer http.ResponseWriter, response *ErrorResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	makeResponse(writer, jsonResponse)
}

func makeResponse(writer http.ResponseWriter, jsonResponse []byte) {
	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, string(jsonResponse))
}
