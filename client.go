package main

import (
	"context"
	"github.com/seogyugim/go-grpc/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial to gRPC Server: %s", err)
	}

	defer conn.Close()

	message := chat.Message{Body: "Hello from the Client"}

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
