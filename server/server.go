package server

import (
	"errors"
	"github.com/3cham/hackathon-chatbot/message"
	"time"
)

type ChatServer struct {
	db ChatDatabase
}

func (s *ChatServer) Register(msg message.RegisterMessage) {
	s.db.RegisterClient(msg.ClientName)
}

func (s *ChatServer) CheckClient(name string) bool {
	return s.db.CheckIfClientRegistered(name)
}

var ErrClientNotAccepted = errors.New("Client is not registered")

func (s *ChatServer) ProcessMessage(msg message.ChatMessage) error {
	if s.CheckClient(msg.ClientName) {
		if msg.Timestamp == "" {
			saveMsg := message.ChatMessage{
				ClientName: msg.ClientName,
				Message: msg.Message,
				Timestamp: time.Time{}.String(),
			}
			s.db.SaveChatMessage(saveMsg)
		} else {
			s.db.SaveChatMessage(msg)
		}
		return nil
	}
	return ErrClientNotAccepted
}

func (s *ChatServer) GetMessageAfter(timestamp string) []message.ChatMessage {
	return s.db.GetMessageAfter(timestamp)
}
