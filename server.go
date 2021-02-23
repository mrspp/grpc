package main

import (
	"context"
	"fmt"
	"grpc/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Hello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	name := request.Name
	response := &hello.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func main() {
	address := "localhost: 9000"

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}
