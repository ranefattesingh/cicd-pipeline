package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ranefattesingh/cicd-pipeline/handler"
)

func TestPingHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	writer := httptest.NewRecorder()
	handler.Ping(writer, req)

	if result := writer.Result(); result.StatusCode != http.StatusOK {
		t.Errorf("Expected status = %v but got status = %v", http.StatusOK, result.StatusCode)
		defer result.Body.Close()
	}
}
