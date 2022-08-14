package json

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func EncodeJSON(writer http.ResponseWriter, data interface{}, statusCode int) {
	byteSlice, err := json.Marshal(data)
	if err != nil {
		EncodeInternalServerError(writer)

		return
	}

	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")

	if _, err := writer.Write(byteSlice); err != nil {
		EncodeInternalServerError(writer)
	}
}

func EncodeResponse(writer http.ResponseWriter, data interface{}, statusCode int) {
	response := Response{
		Success: true,
		Data:    data,
	}

	EncodeJSON(writer, response, statusCode)
}
