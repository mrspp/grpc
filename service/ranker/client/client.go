package main

import (
	"context"
	"fmt"
	rank "grpc/service/ranker/ranking/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from client ranker...")

	clientConn, err := grpc.Dial("localhost: 5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	client := rank.NewRankServiceClient(clientConn)
	request := &rank.RankRequest{Index: "1", ProductName: "suha", Rating: "4"}

	resp, _ := client.RankService(context.Background(), request)

	fmt.Printf("Receive response => [%v]", resp.Popularity)

}
