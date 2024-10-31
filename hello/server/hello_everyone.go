package main

import (
	"io"
	"log"

	pb "grpc-unary/hello/proto"
)

func (*Server) HelloEveryone(stream pb.HelloService_HelloEveryoneServer) error {
	log.Println("HelloEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.HelloResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
