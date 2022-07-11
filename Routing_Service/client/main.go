package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/routingpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Client dialing on port 3000")
	cc, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}

	c := routingpb.NewRoutingServiceClient(cc)
	req := routingpb.AvailabiltyRequest{
		Source: &routingpb.Location{
			Latitude:  2,
			Longitude: 3,
		},
	}
	str, err := c.GetAvailability(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Get Availability RPC: %v", err)
	}
	if err := str.Send(&req); err != nil {
		log.Fatalf("Error while send to Server: %v", err)
	}
	if err := str.CloseSend(); err != nil {
		log.Fatalf("Error while close send: %v", err)
	}
	count := 0
	var a []*routingpb.AvailabiltyResponse
	for {
		res, err := str.Recv()
		if err != nil {
			break
		}
		count++
		log.Printf("Response from Get Availability RPC: %v and %v", res.GetAvailabiltyResponse(), count)
		a = res.GetAvailabiltyResponse()
	}
	for _, ri := range a {
		fmt.Printf("%s\t%s\t%v\n",
			ri.GetCarType(),
			ri.GetLocation(),
			ri.GetDistance())
	}
}
