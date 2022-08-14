package handler

import (
	"net/http"

	"github.com/ranefattesingh/cicd-pipeline/handler/json"
)

func NewPingHandler() *pingHandle {
	return &pingHandle{}
}

type pingHandle struct {
}

func (*pingHandle) Ping(w http.ResponseWriter, r *http.Request) {
	json.EncodeResponse(w, "pong!", http.StatusOK)
}
