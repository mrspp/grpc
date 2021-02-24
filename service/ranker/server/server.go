package server

import (
	"context"
	"fmt"
	rank "grpc/service/ranker/ranking/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Ranker(ctx context.Context, request *rank.RankRequest) (*rank.RankResponse, error) {
	index := request.Index
	productname := request.ProductName
	rating := request.Rating

	response := &rank.RankResponse{
		Popularity: string(index) + productname + rating,
	}
	return response, nil
}

func main() {
	address := "localhost:5000"
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ", address)

	s := grpc.NewServer()

	rank.RegisterRankServiceServer(s, &server{})

	s.serve(lis)
}
