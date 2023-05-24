package chat

import (
	"fmt"
	"os"
	"testing"

	"github.com/gorilla/websocket"
)

func TestWorldChat(t *testing.T) {
	host := os.Getenv("APP_HOST")
	url := fmt.Sprintf("ws://%s/world_chat/1", host)
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("Failed to dial WebSocket: %v", err)
	}
	defer ws.Close()

	message := "Testmessage"
	err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message: %v", err)
	}
	
	received := string(p)

	if received != message {
		t.Errorf("Received message does not match sent message: expected=%q, actual=%q", message, received)
	}
}