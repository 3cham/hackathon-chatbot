package server

import (
	"errors"
	"github.com/3cham/hackathon-chatbot/message"
)

type ChatServer struct {
	Clients []string
	ChatMessages []message.ChatMessage
}

func (s *ChatServer) Register(msg message.RegisterMessage) {
	s.Clients = append(s.Clients, msg.ClientName)
}

func (s *ChatServer) CheckClient(name string) bool {
	result := false
	for _, savedClient := range(s.Clients) {
		if savedClient == name {
			result = true
			break
		}
	}
	return result
}

var ErrClientNotAccepted = errors.New("Client is not registered")

func (s *ChatServer) ProcessMessage(msg message.ChatMessage) error {
	if s.CheckClient(msg.ClientName) {
		s.ChatMessages = append(s.ChatMessages, msg)
		return nil
	}
	return ErrClientNotAccepted
}

func (s *ChatServer) GetMessageAfter(timestamp string) []message.ChatMessage {
	result := []message.ChatMessage{}
	for _, msg := range s.ChatMessages {
		if msg.Timestamp > timestamp {
			result = append(result, msg)
		}
	}
	return result
}
