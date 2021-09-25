package main

import (
	"fmt"
	"learn_gRPC/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Hello Worlds")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}

}
