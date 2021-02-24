package main

import (
	"context"
	"fmt"
	auth "grpc/service/authentication/auth/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client authentication")

	otps := grpc.WithInsecure()

	clientConn, err := grpc.Dial("localhost:8000", otps)
	if err != nil {
		log.Fatal(err)
	}

	defer clientConn.Close()

	client := auth.NewAuthServiceClient(clientConn)
	request := &auth.AuthRequest{Username: "phuong", Password: "123456"}

	resp, _ := client.Auth(context.Background(), request)

	fmt.Printf("Receive response => [%v, %v]", resp.Mes, resp.Token)
}
