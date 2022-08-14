package json

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func EncodeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	byteSlice, err := json.Marshal(data)
	if err != nil {
		EncodeInternalServerError(w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(byteSlice)
	if err != nil {
		EncodeInternalServerError(w)
	}
}

func EncodeResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	response := Response{
		Success: true,
		Data:    data,
	}

	EncodeJSON(w, response, statusCode)
}
