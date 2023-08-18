package main

import (
	"context"
	pb "github.com/Powehi-cs/Go-Grpc/hello"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "127.0.0.1:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection error: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloServerClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Hello World!",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("[*] get server greeting response: %v", r.Message)
}
