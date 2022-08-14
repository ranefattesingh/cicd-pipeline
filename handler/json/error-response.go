package json

import (
	"errors"
	"net/http"
)

var (
	errInternalServer = errors.New("server could not handle the request")
)

func EncodeError(w http.ResponseWriter, data interface{}, statusCode int) {
	res := Response{
		Success: false,
		Data:    data,
	}

	EncodeJSON(w, res, statusCode)
}

func EncodeInternalServerError(w http.ResponseWriter) {
	EncodeError(w, errInternalServer, http.StatusInternalServerError)
}
