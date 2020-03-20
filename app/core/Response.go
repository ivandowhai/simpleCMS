package core

import (
	"io"
	"net/http"
)

func MakeResponse(writer http.ResponseWriter, data []byte) {
	var settings = GetSettings()
	writer.Header().Set("Content-type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", settings.FrontHost)
	writer.Header().Set("Access-Control-Request-Method", "GET, POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	io.WriteString(writer, string(data))
}
