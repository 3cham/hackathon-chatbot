package server

import (
	"github.com/3cham/hackathon-chatbot/message"
	"testing"
)

func TestServerRegister(t *testing.T) {
	t.Run("Server should accept register from client", func(t *testing.T) {
		server := &ChatServer{}
		clientName := "TestClient"
		invalidClientName := "InvalidClient"

		msg := &message.RegisterMessage{
			ClientName: clientName,
		}
		server.Register(msg)
		if !server.CheckClient(clientName) {
			t.Fatalf("Server has problem. Client with name %s should be found.", clientName)
		}
		if server.CheckClient(invalidClientName) {
			t.Fatalf("Server has problem. Client with name %s should not be found", invalidClientName)
		}
	})
}

func TestServerSendMessage(t *testing.T) {
	t.Run("Server should accept message sent from client", func(t *testing.T) {
		server := &ChatServer{}
		clientName := "ValidClient"
		registerMessage := &message.RegisterMessage{
			ClientName: clientName,
		}
		server.Register(registerMessage)

		msg := message.ChatMessage{
			ClientName: clientName,
			Message: "Hello Server",
		}

		err := server.ProcessMessage(msg)
		if err != nil {
			t.Fatalf("Error is not expected because client has been registered. Got: %v", err)
		}

		if len(server.ChatMessages) == 0 {
			t.Fatalf("Message is currently not saved inside the server")
		}
	})

	t.Run("Server should not accept message sent from invalid client", func(t *testing.T) {
		server := &ChatServer{}
		clientName := "invalidClient"
		msg := message.ChatMessage{
			ClientName: clientName,
			Message: "Hello Server",
		}

		err := server.ProcessMessage(msg)
		if err != ErrClientNotAccepted {
			t.Fatalf("Server should not accept message from client. Got: %v", err)
		}
	})
}
