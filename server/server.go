package server

import (
	"encoding/json"
	"errors"
	"github.com/3cham/hackathon-chatbot/message"
	"net/http"
	"net/url"
	"time"
)

type ChatServer struct {
	db ChatDatabase
}

func (s *ChatServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.String() == "/api/register" {
		s.HandleRegister(writer, request)
	} else
	if request.URL.String() == "/api/send_message" {
		s.HandleSendMessage(writer, request)
	} else
	if request.URL.Path == "/api/get_messages" {
		s.HandleGetMessage(writer, request)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
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
				Timestamp: time.Now().String(),
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

func (s *ChatServer) HandleRegister(writer http.ResponseWriter, request *http.Request) {
	msg := &message.RegisterMessage{}
	if err := json.NewDecoder(request.Body).Decode(msg); err == nil && msg.ClientName != "" {
		s.Register(*msg)
		writer.WriteHeader(http.StatusAccepted)
		return
	}
	writer.WriteHeader(http.StatusNotAcceptable)
}

func (s *ChatServer) HandleSendMessage(writer http.ResponseWriter, request *http.Request) {
	msg := &message.ChatMessage{}
	if err := json.NewDecoder(request.Body).Decode(msg); err == nil && msg.ClientName != "" && msg.Message != "" {
		if err := s.ProcessMessage(*msg); err == nil {
			writer.WriteHeader(http.StatusAccepted)
			return
		}
	}
	writer.WriteHeader(http.StatusNotAcceptable)
}

func (s *ChatServer) HandleGetMessage(writer http.ResponseWriter, request *http.Request) {
	if u, err := url.ParseRequestURI(request.URL.String()); err == nil {
		if query, err := url.ParseQuery(u.RawQuery); err == nil {
			fromTimestamp, successful := query["from"]
			if successful {
				result := s.GetMessageAfter(fromTimestamp[0])
				json.NewEncoder(writer).Encode(result)
				writer.WriteHeader(http.StatusOK)
				return
			}
		}
	}
	writer.WriteHeader(http.StatusBadRequest)
}

func NewChatServer() *ChatServer {
	return &ChatServer{&InmemoryDatabase{}}
}
