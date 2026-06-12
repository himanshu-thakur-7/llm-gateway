package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	gatewayerrors "github.com/himanshu-thakur-7/llm-gateway/gatewayerrors"
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
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
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

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.groq.com/openai/v1/chat/completions",
		bytes.NewBuffer(body),
	)
	if err != nil {
		fmt.Printf("Error occured %v \n", err)
		return types.ChatCompletionResponse{}, err
	}

	httpReq.Header.Set(
		"Authorization",
		"Bearer "+g.apiKey,
	)

	httpReq.Header.Set(
		"Content-Type",
		"application/json",
	)

	fmt.Println(httpReq)

	resp, err := g.client.Do(httpReq)

	if err != nil {
		fmt.Printf("Error occured %v \n", err)
		return types.ChatCompletionResponse{}, err
	}

	defer resp.Body.Close()

	fmt.Printf("response %s", resp.Status)

	// bodyBytes, err := io.ReadAll(resp.Body)

	// if err != nil {
	// 	return types.ChatCompletionResponse{}, err
	// }

	// fmt.Println(string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		var errType gatewayerrors.ErrorType
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			errType = gatewayerrors.ErrorTypeAuth
		case http.StatusTooManyRequests:
			errType = gatewayerrors.ErrorTypeRateLimit
		default:
			errType = gatewayerrors.ErrorTypeProvider
		}
		return types.ChatCompletionResponse{},
			&gatewayerrors.ProviderError{
				Type:     errType,
				Provider: g.Name(),
				Message: fmt.Sprintf(
					"provider returned status %d", resp.StatusCode,
				),
			}
	}

	var groqResp groqResponse

	if err := json.NewDecoder(resp.Body).Decode(&groqResp); err != nil {
		fmt.Printf("Error occured line 131%v \n", err)
		return types.ChatCompletionResponse{}, &gatewayerrors.ProviderError{
			Type:     gatewayerrors.ErrorTypeTimeout,
			Provider: g.Name(),
			Message:  err.Error(),
		}
	}

	if len(groqResp.Choices) == 0 {
		return types.ChatCompletionResponse{},
			fmt.Errorf("groq returned no choices")
	}

	return types.ChatCompletionResponse{
		ID:      groqResp.ID,
		Content: groqResp.Choices[0].Message.Content,
	}, nil
}

func (g *GroqProvider) SupportsModel(
	model string,
) bool {
	_, ok := g.supportedModels[model]
	return ok
}
