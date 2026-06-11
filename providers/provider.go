package providers

import (
	"context"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type Provider interface {
	Name() string
	SupportsModel(model string) bool
	ChatCompletion(ctx context.Context, req types.ChatCompletionRequest) (types.ChatCompletionResponse, error)
}
