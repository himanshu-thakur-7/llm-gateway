package main

import (
	"log"
	"net/http"

	"github.com/himanshu-thakur-7/llm-gateway/endpoints"
	"github.com/himanshu-thakur-7/llm-gateway/providers"
	"github.com/himanshu-thakur-7/llm-gateway/router"
)

func main() {
	http.HandleFunc("/healthz", endpoints.HealthzHandler)

	mockProvider := providers.MockProvider{}
	anthropicProvider := providers.AnthropicProvider{}
	providerRegistry := map[string]providers.Provider{
		"mock":      mockProvider,
		"anthropic": anthropicProvider,
	}
	gatewayRouter := router.NewStaticRouter(providerRegistry)

	http.HandleFunc("/v1/chat/completions", endpoints.NewChatCompletionHandler(gatewayRouter))

	log.Println("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
