package router

import (
	"context"
	"fmt"

	"github.com/himanshu-thakur-7/llm-gateway/providers"
	"github.com/himanshu-thakur-7/llm-gateway/registry"
	"github.com/himanshu-thakur-7/llm-gateway/types"
)

type StaticRouter struct {
	providers map[string]providers.Provider
}

func NewStaticRouter(
	providers map[string]providers.Provider,
) *StaticRouter {
	return &StaticRouter{
		providers: providers,
	}
}

func (r *StaticRouter) Route(
	ctx context.Context,
	req types.ChatCompletionRequest,
) (types.ChatCompletionResponse, error) {
	providerName, ok := registry.ModelRegistry[req.Model]
	if !ok {
		return types.ChatCompletionResponse{},
			fmt.Errorf("unknown model: %s", req.Model)
	}
	provider, ok := r.providers[providerName]

	if !ok {
		return types.ChatCompletionResponse{}, fmt.Errorf("provider not found: %s", providerName)
	}

	return provider.ChatCompletion(
		ctx, req,
	)
}
