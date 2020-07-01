package main

import (
	"fmt"
	"go-grpc/ws"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var wsServer = ws.New()

	var grpcServer = grpc.NewServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ws.RegisterWebsocketServer(grpcServer, wsServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}

	}()

	fmt.Println("serve echo")
	wsServer.Start(":8000")
}
