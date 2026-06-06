package types

import (
	"errors"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func (r ChatCompletionRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}

	if len(r.Messages) == 0 {
		return errors.New("atleast one message is required")
	}

	return nil
}
