package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/himanshu-thakur-7/llm-gateway/router"
	"github.com/himanshu-thakur-7/llm-gateway/types"
)

func NewChatCompletionHandler(router router.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			log.Println(err)
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
		resp, _ := router.Route(
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
}
