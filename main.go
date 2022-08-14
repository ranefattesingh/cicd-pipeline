package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranefattesingh/cicd-pipeline/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping)
	log.Println("Server starting on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server error")
	}
}
