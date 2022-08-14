package handler

import (
	"net/http"

	"github.com/ranefattesingh/cicd-pipeline/handler/json"
)

func Ping(writer http.ResponseWriter, _ *http.Request) {
	json.EncodeResponse(writer, "pong!", http.StatusOK)
}
