// routes/openai.go

package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

// Initialize the OpenAI client
var openaiClient *openai.Client

func InitOpenAIClient(client *openai.Client) {
	openaiClient = client
}

// handleOpenAI handles requests to interact with OpenAI
func HandleOpenAI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Interact with OpenAI
	completion, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.Message,
				},
			},
		},
	)
	if err != nil {
		log.Printf("ChatCompletion error: %v", err)
		http.Error(w, "Error interacting with OpenAI", http.StatusInternalServerError)
		return
	}

	// Respond with OpenAI's response
	response := struct {
		Response string `json:"response"`
	}{
		Response: completion.Choices[0].Message.Content,
	}
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
