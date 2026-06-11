package providers

import (
	"context"
	"fmt"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type AnthropicProvider struct{}

func (m AnthropicProvider) ChatCompletion(
	ctx context.Context,
	req types.ChatCompletionRequest,
) (types.ChatCompletionResponse, error) {
	return types.ChatCompletionResponse{
		ID: "Anthropic-123",
		Content: fmt.Sprintf(
			"response from provider %s",
			m.Name(),
		),
	}, nil
}

func (m AnthropicProvider) Name() string {
	return "anthropic"
}

func (m AnthropicProvider) SupportsModel(model string) bool {
	return model == "claude"
}
