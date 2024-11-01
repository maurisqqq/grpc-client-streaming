package main

import (
	"context"
	"log"
	"time"

	pb "grpc-client-streaming/hello/proto"
)

func doLongHello(c pb.HelloServiceClient) {
	log.Println("doLongHello was invoked")

	reqs := []*pb.HelloRequest{
		{FirstName: "Maulana"},
		{FirstName: "Risqi"},
		{FirstName: "Mustofa"},
	}

	stream, err := c.LongHello(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongHello: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongHello: %v\n", err)
	}

	log.Printf("LongHello: %s\n", res.Result)
}
