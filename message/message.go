package message

type RegisterMessage struct {
	ClientName string
}

type ChatMessage struct {
	ClientName string
	Message string
	Timestamp string
}