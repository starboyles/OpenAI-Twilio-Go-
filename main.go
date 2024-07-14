package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/sashabaranov/go-openai"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	// Read environment variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	openaiKey := os.Getenv("OPENAI_API_KEY")

	// Check if environment variables are set
	if accountSid == "" || authToken == "" || openaiKey == "" {
		log.Fatal("TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, and OPENAI_API_KEY must be set")
	}

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Use the Twilio client...
	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+233200356369")
	params.SetFrom("+15855492374")
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatal("Error sending SMS:", err)
	}

	fmt.Println("Message SID:", *resp.Sid)

	// Initialize OpenAI client
	openaiClient := openai.NewClient(openaiKey)

	// Use OpenAI client
	completion, err := openaiClient.CreateChatCompletion(
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
		log.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println("OpenAI Response:", completion.Choices[0].Message.Content)
}