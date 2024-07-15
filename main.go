// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"your-module-path/controllers"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/send-sms", controllers.SendSMS)
	router.POST("/chat", controllers.Chat)

	// Start server
	router.Run(":8080")
}
