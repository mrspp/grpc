package main

import (
	"context"
	"fmt"
	"grpc/hello"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	otps := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:9000", otps)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := hello.NewHelloServiceClient(cc)
	request := &hello.HelloRequest{Name: "Jenny"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}
