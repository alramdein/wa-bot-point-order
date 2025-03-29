package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Define message structure
type ButtonMessage struct {
	To      string `json:"to"`
	MsgType string `json:"msgType"`
	Body    string `json:"body"`
	Buttons []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"buttons"`
}

func sendWhatsAppMessage() {
	// UltraMsg API Endpoint
	apiURL := "https://api.ultramsg.com/YOUR_INSTANCE_ID/messages/chat"

	// Create the message with buttons
	message := ButtonMessage{
		To:      "6281234567890", // Replace with recipient number
		MsgType: "interactive",
		Body:    "Choose an option below:",
		Buttons: []struct {
			ID   string `json:"id"`
			Text string `json:"text"`
		}{
			{ID: "order_laundry", Text: "🧺 Order Laundry"},
			{ID: "make_payment", Text: "💰 Make Payment"},
		},
	}

	// Convert message to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send request to UltraMsg API
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "YOUR_API_KEY") // Replace with your UltraMsg API key

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Message sent successfully!")
}

func main() {
	sendWhatsAppMessage()
}
