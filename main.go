// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/starboyles/twilio-gemini-assistant/controllers"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/voice", controllers.HandleIncomingCall)
	router.POST("/gather", controllers.GatherSpeech)
	router.POST("/process-speech", controllers.ProcessSpeech)

	// Start server
	router.Run(":8080")
}