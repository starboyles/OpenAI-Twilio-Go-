// controllers/chat_controller.go
package controllers

import (
	"context"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// Chat handles the interaction with OpenAI API
func Chat(c *gin.Context) {
	// Read environment variable
	openaiKey := os.Getenv("OPENAI_API_KEY")

	if openaiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API key must be set"})
		return
	}

	// Initialize OpenAI client
	client := openai.NewClient(openaiKey)

	// Create chat completion
	completion, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello, how are you?",
				},
			},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": completion.Choices[0].Message.Content})
}
