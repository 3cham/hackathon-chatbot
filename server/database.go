package server

import "github.com/3cham/hackathon-chatbot/message"

type ChatDatabase interface {
	RegisterClient(name string)
	CheckIfClientRegistered(name string) bool
	SaveChatMessage(msg message.ChatMessage)
	GetMessageAfter(timestamp string) []message.ChatMessage
}

type InmemoryDatabase struct {
	Clients []string
	ChatMessages []message.ChatMessage
}

func (d *InmemoryDatabase) RegisterClient(name string) {
	d.Clients = append(d.Clients, name)
}

func (d *InmemoryDatabase) CheckIfClientRegistered(name string) bool {
	result := false
	for _, savedClient := range d.Clients {
		if savedClient == name {
			result = true
			break
		}
	}
	return result
}

func (d *InmemoryDatabase) SaveChatMessage(msg message.ChatMessage) {
	d.ChatMessages = append(d.ChatMessages, msg)
}

func (d *InmemoryDatabase) GetMessageAfter(timestamp string) []message.ChatMessage {
	result := []message.ChatMessage{}
	for _, msg := range d.ChatMessages {
		if msg.Timestamp > timestamp {
			result = append(result, msg)
		}
	}
	return result
}
