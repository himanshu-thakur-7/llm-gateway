package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/himanshu-thakur-7/llm-gateway/providers"
	"github.com/himanshu-thakur-7/llm-gateway/types"
)

func ChatCompletionHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	// Decode request body
	var req types.ChatCompletionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	// Mock response
	resp, _ := providers.MockProvider{}.ChatCompletion(
		r.Context(), req,
	)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(
			w,
			"failed to encode response",
			http.StatusInternalServerError,
		)
		return
	}
}
