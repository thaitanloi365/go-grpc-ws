package main

import (
	"context"
	"go-grpc/ws"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := ws.NewWebsocketClient(conn)

	response, err := c.SendMessage(context.Background(), &ws.Message{UserId: "asdfasdf"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.GetUserId())
}
