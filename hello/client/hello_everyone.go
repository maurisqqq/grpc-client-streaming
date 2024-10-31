package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-unary/hello/proto"
)

func doHelloEveryone(c pb.HelloServiceClient) {
	log.Println("doHelloEveryone was invoked")

	stream, err := c.HelloEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	requests := []*pb.HelloRequest{
		{FirstName: "Maulana"},
		{FirstName: "Risqi"},
		{FirstName: "Mustofa"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc
}
