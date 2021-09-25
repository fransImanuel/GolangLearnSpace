package main

import (
	"fmt"
	"learn_gRPC/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect:%v", err)
	}
	defer cc.Close()

	c := greetpb.File_greet_greetpb_greet_proto
	fmt.Printf("Created client: %f", c)
}
