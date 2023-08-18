package main

import (
	"context"
	pb "github.com/Powehi-cs/Go-Grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type HelloServer struct {
	pb.UnimplementedHelloServerServer
}

func (hs *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (out *pb.HelloReply, err error) {
	return &pb.HelloReply{
		Message: "hello grpc ===> " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServerServer(s, &HelloServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: =%v", err)
	}
}
