package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type GroqProvider struct {
	apiKey string
	client *http.Client

	supportedModels map[string]bool
}

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model    string        `json:"model"`
	Messages []groqMessage `json:"messages"`
}

type groqResponse struct {
	ID string `json:"id"`

	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewGroqProvider(apiKey string, supportedModels map[string]bool) *GroqProvider {
	return &GroqProvider{
		apiKey:          apiKey,
		client:          &http.Client{},
		supportedModels: supportedModels,
	}
}

func (g *GroqProvider) Name() string {
	return "groq"
}

func (g *GroqProvider) ChatCompletion(
	ctx context.Context,
	req types.ChatCompletionRequest,
) (types.ChatCompletionResponse, error) {
	fmt.Println(req)
	// if !g.SupportsModel(req.Model) {
	// 	return types.ChatCompletionResponse{},
	// 		fmt.Errorf(
	// 			"provider %s does not support model %s",
	// 			g.Name(),
	// 			req.Model,
	// 		)
	// }

	messages := make([]groqMessage, 0, len(req.Messages))

	for _, msg := range req.Messages {
		messages = append(messages, groqMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	groqRequest := groqRequest{
		Model:    req.Model,
		Messages: messages,
	}
	body, err := json.Marshal(groqRequest)
	if err != nil {
		return types.ChatCompletionResponse{}, err
	}

	fmt.Printf("%s\n", string(body))

	return types.ChatCompletionResponse{}, nil
}

func (g *GroqProvider) SupportsModel(
	model string,
) bool {
	_, ok := g.supportedModels[model]
	return ok
}
