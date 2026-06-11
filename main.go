package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/himanshu-thakur-7/llm-gateway/endpoints"
	"github.com/himanshu-thakur-7/llm-gateway/providers"
	"github.com/himanshu-thakur-7/llm-gateway/router"
)

func main() {
	http.HandleFunc("/healthz", endpoints.HealthzHandler)

	mockProvider := providers.MockProvider{}
	anthropicProvider := providers.AnthropicProvider{}
	groqProvider := providers.NewGroqProvider(
		os.Getenv("GROQ_API_KEY"), map[string]bool{"llama-3.3-70b-versatile": true})

	providerRegistry := map[string]providers.Provider{
		"mock":      mockProvider,
		"anthropic": anthropicProvider,
		"groq":      groqProvider,
	}

	fmt.Println(providerRegistry)
	gatewayRouter := router.NewStaticRouter(providerRegistry)

	http.HandleFunc("/v1/chat/completions", endpoints.NewChatCompletionHandler(gatewayRouter))

	log.Println("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
