package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/3cham/hackathon-chatbot/message"
	"log"
	"net/http"
	"os"
	"time"
)

const ContentType = "application/json"

type ChatClient struct {
	Name string
	ServerAddress string
	client http.Client
}

func (c ChatClient) Start() {
	// Implement input waiting & message getting from server
	c.Register()
	go c.ReceiveMessage()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		c.SendMessage(text)
	}
}

func (c ChatClient) Register() {
	registerMessage := message.RegisterMessage{
		ClientName: c.Name,
	}

	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(registerMessage)

	c.client.Post(c.ServerAddress + "/api/register", ContentType, body)
}

func (c ChatClient) ReceiveMessage() {
	lastTimestamp := "0"
	for {
		time.Sleep(5 * time.Second)

		newMessages := c.GetMessage(lastTimestamp)
		if len(newMessages) > 0 {
			lastTimestamp = c.printOut(newMessages)
		}
	}
}

func (c ChatClient) SendMessage(text string) {
	chatMessage := message.ChatMessage{
		ClientName: c.Name,
		Message: text,
	}

	body := &bytes.Buffer{}
	json.NewEncoder(body).Encode(chatMessage)

	c.client.Post(c.ServerAddress + "/api/send_message", ContentType, body)
}

func (c ChatClient) GetMessage(timestamp string) []message.ChatMessage{

	response, _ := c.client.Get(c.ServerAddress + "/api/get_messages?from=" + timestamp)
	messages := &[]message.ChatMessage{}

	json.NewDecoder(response.Body).Decode(messages)
	return *messages
}

func (c ChatClient) printOut(messages []message.ChatMessage) string {
	newTimestamp := ""
	for _, msg := range messages {
		if msg.ClientName != c.Name {
			fmt.Printf("%s: %s", msg.ClientName, msg.Message)
		}
		if msg.Timestamp > newTimestamp {
			newTimestamp = msg.Timestamp
		}
	}
	return newTimestamp
}

func NewChatClient(serverAdd string, name string) {

	log.Printf("Client with name %s is started against server %s", name, serverAdd)
	client := ChatClient{Name: name, ServerAddress: serverAdd, 	client: http.Client{}}
	client.Start()
}