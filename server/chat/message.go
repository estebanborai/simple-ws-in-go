package chat

import "time"

// Message represents a chat message
type Message struct {
	Author   *Client
	Message  string
	IssuedAt time.Time
}

// NewMessage creates an instance of Message and returns
// a pointer
func NewMessage(author *Client, message string) *Message {
	var msg *Message = new(Message)

	msg.Author = author
	msg.Message = message
	msg.IssuedAt = time.Now()

	return msg
}
