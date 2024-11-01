package main

import (
	"log"
	"net"

	pb "grpc-client-streaming/hello/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.HelloServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
