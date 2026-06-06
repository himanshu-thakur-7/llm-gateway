package providers

import (
	"context"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type MockProvider struct{}

func (m MockProvider) ChatCompletion(
	ctx context.Context,
	req types.ChatCompletionRequest,
) (types.ChatCompletionResponse, error) {
	return types.ChatCompletionResponse{
		ID:      "mock-123",
		Content: "Hello from Mock Provider",
	}, nil
}
