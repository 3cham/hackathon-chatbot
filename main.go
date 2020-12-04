package main

import (
	"flag"
	"github.com/3cham/hackathon-chatbot/client"
	"github.com/3cham/hackathon-chatbot/server"
)

func main()  {

	typePtr := flag.String("type", "`server` or `client`", "Start server with client")
	addPtr := flag.String("address", "server ip address (with port)", "Tell client where the server is")
	namePtr := flag.String("name", "name of the chat client", "To differentiate chat member")
	flag.Parse()

	if *typePtr == "server" {
		server.NewChatServer()
	} else
	if *addPtr != "" && *namePtr != "" {
		client.NewChatClient(*addPtr, *namePtr)
	}
}