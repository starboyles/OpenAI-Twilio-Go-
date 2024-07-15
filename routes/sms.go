// routes/sms.go

package routes

import (
	"fmt"
	"log"
	"net/http"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/twilio/twilio-go"
)

// Initialize the Twilio client
var twilioClient *twilio.RestClient

func InitTwilioClient(client *twilio.RestClient) {
	twilioClient = client
}

// handleSMS handles incoming SMS messages from Twilio
func HandleSMS(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse incoming SMS message
	from := r.FormValue("From")
	body := r.FormValue("Body")
	log.Printf("Received SMS from %s: %s", from, body)

	// Respond with a message
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(from)
	params.SetFrom("+15855492374")
	params.SetBody("Thank you for your message!")

	resp, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		http.Error(w, "Error sending SMS", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message SID: %s", *resp.Sid)
}
