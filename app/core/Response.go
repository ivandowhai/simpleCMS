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

func makeResponse(writer http.ResponseWriter, data []byte) {
	var settings = GetSettings()
	writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", settings.FrontHost)
	writer.Header().Set("Access-Control-Max-Age", "86400")
	writer.Header().Set("Access-Control-Request-Method", "GET, POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Type,access-control-allow-origin, access-control-allow-headers")
	writer.WriteHeader(http.StatusOK)

	io.WriteString(writer, string(data))
}
