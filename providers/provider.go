package providers

import (
	"context"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type Provider interface {
	ChatCompletion(ctx context.Context, req types.ChatCompletionRequest) (types.ChatCompletionResponse, error)
}
