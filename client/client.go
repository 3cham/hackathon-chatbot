package client

type ChatClient struct {
	ServerAddress string
}

func (c ChatClient) Start() {
	// Implement input waiting & message getting from server
}

func NewChatClient(serverAdd string) ChatClient {
	return ChatClient{serverAdd}
}