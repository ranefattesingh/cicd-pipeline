package json

import (
	"errors"
	"net/http"
)

var errInternalServer = errors.New("server could not handle the request")

func EncodeError(writer http.ResponseWriter, data interface{}, statusCode int) {
	res := Response{
		Success: false,
		Data:    data,
	}

	EncodeJSON(writer, res, statusCode)
}

func EncodeInternalServerError(writer http.ResponseWriter) {
	EncodeError(writer, errInternalServer, http.StatusInternalServerError)
}
