package main

import (
	"github.com/3cham/hackathon-chatbot/client"
	"github.com/3cham/hackathon-chatbot/server"
	"log"
	"net/http"
	"flag"
)

func main()  {

	typePtr := flag.String("type", "`server` or `client`", "Start server with client")
	addPtr := flag.String("address", "server ip address (with port)", "Tell client where the server is")

	flag.Parse()
	if *typePtr == "server" {
		srv := server.NewChatServer()

		if err := http.ListenAndServe(":5000", srv); err != nil {
			log.Fatalf("could not listen on port 5000 %v", err)
		}
	} else {
		client := client.NewChatClient(*addPtr)
		client.Start()
	}
}