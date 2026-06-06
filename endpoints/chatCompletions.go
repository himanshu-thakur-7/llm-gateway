package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ChatCompletionRequest struct {
	Model   string `json:"model"`
	Message string `json:"message"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (r ChatCompletionRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}

	if r.Message == "" {
		return errors.New("message is required")
	}

	return nil
}

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
	var req ChatCompletionRequest

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
	resp := ChatCompletionResponse{
		ID:      "mock-123",
		Content: "Hello from gateway",
	}

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
