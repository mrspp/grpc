package main

import (
	"context"
	"fmt"
	auth "grpc/service/authentication/auth/proto"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Auth(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {
	username := request.Username
	password := request.Password
	response := &auth.AuthResponse{
		Mes:   "From" + username + password,
		Token: "Token: xxxxx",
	}
	return response, nil
}

func main() {
	address := "localhost: 8000"

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &server{})

	s.Serve(lis)
}
