package chat

import (
	"context"
	"log"
)

//Server ...
type Server struct {
}

// SayHello ...
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello from the server"}, nil
}

//BroadcastMessage /
func (s *Server) BroadcastMessage(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Broadcasting new message from a client: %s", in.Body)
	return &Message{Body: "Broadcasted message!"}, nil
}
