package main

import (
	"github.com/3cham/hackathon-chatbot/server"
	"log"
	"net/http"
)

func main()  {
	srv := server.NewChatServer()

	if err := http.ListenAndServe(":5000", srv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}