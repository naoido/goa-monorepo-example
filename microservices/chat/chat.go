package chatapi

import (
	"context"
	chat "goa-example/microservices/chat/gen/chat"

	"goa.design/clue/log"
)

// chat service example implementation.
// The example methods log the requests and return zero values.
type chatsrvc struct{}

// NewChat returns the chat service implementation.
func NewChat() chat.Service {
	return &chatsrvc{}
}

// Creates a new chat room.
func (s *chatsrvc) CreatRoom(ctx context.Context) (res string, err error) {
	log.Printf(ctx, "chat.creat-room")
	return
}
