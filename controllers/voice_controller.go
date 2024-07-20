// controllers/voice_controller.go
package controllers

import (
	"log"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/twilio/twilio-go"
)


func HandleIncomingCall(c *gin.Context) {
	response := twiml.NewVoiceResponse()
	gather := response.Gather(twiml.Gather{
		Input:  "speech",
		Action: "/gather",
	})
	gather.Say("Hello, I am your personal assistant. How can I help you today?")

	c.XML(http.StatusOK, response)
}

func GatherSpeech(c *gin.Context) {
	speechResult := c.PostForm("SpeechResult")
	if speechResult == "" {
		log.Println("No speech result received")
		c.String(http.StatusBadRequest, "No speech result received")
		return
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Println("OpenAI API key not found in environment")
		c.String(http.StatusInternalServerError, "OpenAI API key not configured")
		return
	}

	// Initialize OpenAI client
	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(
		c,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: speechResult,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		c.String(http.StatusInternalServerError, "Error processing speech")
		return
	}

	aiResponse := resp.Choices[0].Message.Content

	response := twiml.NewVoiceResponse()
	response.Say(aiResponse)
	gather := response.Gather(twiml.Gather{
		Input:  "speech",
		Action: "/gather",
	})
	gather.Say("Do you have any other questions?")

	c.XML(http.StatusOK, response)
}