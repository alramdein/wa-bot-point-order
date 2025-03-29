package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WhatsAppWebhook struct {
	From string `json:"from"`
	Text struct {
		Body string `json:"body"`
	} `json:"text"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON request
	var message WhatsAppWebhook
	if err := json.Unmarshal(body, &message); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Handle user response
	switch message.Text.Body {
	case "order_laundry":
		fmt.Println("User wants to order laundry!")
		// Send response message (can integrate with DB)
	case "make_payment":
		fmt.Println("User wants to make payment!")
		// Generate payment link and send it
	default:
		fmt.Println("Unknown command:", message.Text.Body)
	}
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
