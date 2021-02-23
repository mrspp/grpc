package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server ...
type Server struct {
}

//SayHello ...
func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
