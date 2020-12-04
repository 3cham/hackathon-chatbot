package server

import (
	"github.com/3cham/hackathon-chatbot/message"
	"testing"
)

func TestServerRegister(t *testing.T) {
	t.Run("Server should accept register from client", func(t *testing.T) {
		server := &ChatServer{
			db: &InmemoryDatabase{},
		}
		clientName := "TestClient"
		invalidClientName := "InvalidClient"

		msg := message.RegisterMessage{
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
		server := &ChatServer{
			db: &InmemoryDatabase{},
		}
		clientName := "ValidClient"
		registerMessage := message.RegisterMessage{
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

		if len(server.GetMessageAfter("0")) == 0 {
			t.Fatalf("Message is currently not saved inside the server")
		}
	})

	t.Run("Server should not accept message sent from invalid client", func(t *testing.T) {
		server := &ChatServer{
			db: &InmemoryDatabase{},
		}
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

func TestGetMessage(t *testing.T)  {
	t.Run("Server should return all messages after a timestamp", func(t *testing.T) {
		server := ChatServer{db: &InmemoryDatabase{}}
		client := "ClientName"
		registerMessage := message.RegisterMessage{
			ClientName: client,
		}

		server.Register(registerMessage)

		m := "Hello Server"
		msg := message.ChatMessage{
			ClientName: client,
			Message:    m,
		}
		server.ProcessMessage(msg)

		msg2 := message.ChatMessage{
			ClientName: client,
			Message:    m,
		}
		server.ProcessMessage(msg2)

		messages := server.GetMessageAfter("0")
		if len(messages) != 2 {
			t.Fatalf("Server does not return messages correctly. Expected 2 message, got %v", messages)
		}

		timestamp := messages[0].Timestamp
		messages = server.GetMessageAfter(timestamp)
		if len(messages) != 1 {
			t.Fatalf("Server does not return messages correctly. Expected 1 message, got %v", messages)
		}
	})
}