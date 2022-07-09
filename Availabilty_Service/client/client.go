package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/availabilitypb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to Dial:$%v", err)
	}
	fmt.Println("Client connection establised...")
	s := availabilitypb.NewAvailabiltyServiceClient(conn)

	req := availabilitypb.AvailabiltyRequest{
		Source: &availabilitypb.Location{
			Latitude:  2,
			Longitude: 3,
		},
	}
	res, err := s.GetAvailability(context.Background(), &req)
	if err != nil {
		log.Fatalf("error occured while calling Get Availability of car:%v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error occured while reading stream %v", err)
		}
		log.Printf("receive:%v\n", msg)
	}
}
