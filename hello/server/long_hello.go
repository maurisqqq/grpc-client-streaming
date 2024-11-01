package main

import (
	"io"
	"log"

	pb "grpc-client-streaming/hello/proto"
)

func (*Server) LongHello(stream pb.HelloService_LongHelloServer) error {
	log.Println("LongHello was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)
		res += "Hello " + req.FirstName + "!\n"
	}
}
