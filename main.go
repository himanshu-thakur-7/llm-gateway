package main

import (
	"log"
	"net/http"

	"github.com/himanshu-thakur-7/llm-gateway/endpoints"
)

func main() {
	http.HandleFunc("/healthz", endpoints.HealthzHandler)

	http.HandleFunc("/v1/chat/completions", endpoints.ChatCompletionHandler)

	log.Println("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
