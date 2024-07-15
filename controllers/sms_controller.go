// controllers/sms_controller.go
package controllers

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SendSMS sends an SMS using Twilio API
func SendSMS(c *gin.Context) {
	// Read environment variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSid == "" || authToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Twilio credentials must be set"})
		return
	}

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Define message parameters
	params := &twilioApi.CreateMessageParams{}
	params.SetTo("")
	params.SetFrom("")
	params.SetBody("Hello from Go!")

	// Send message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println("Error sending SMS:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message_sid": *resp.Sid})
}
