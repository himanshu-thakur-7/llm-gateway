package router

import (
	"context"

	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type Router interface {
	Route(
		ctx context.Context,
		req types.ChatCompletionRequest,
	) (types.ChatCompletionResponse, error)
}
