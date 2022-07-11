package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/availabilitypb"
	"github.com/Dinesh789kumar12/POC1/CarAppPb/routingpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Car struct {
	CarType  string `json:"car"`
	Location string `json:"location"`
	Distance int32  `json:"distance"`
}

type AvailableResponse struct {
	listAvailableCar []Car
}

func main() {
	router := gin.Default()
	rg := router.Group("api/v1/carapp")
	{
		rg.GET("/car", fetchAvailableCarNearby)
		rg.GET("/car1", fetchAvailableCarNear)
	}

	router.Run()
}

func fetchAvailableCarNearby(c *gin.Context) {

	var carpool []Car
	sAddress := ":50052"
	conn, e := grpc.Dial(sAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := availabilitypb.NewAvailabiltyServiceClient(conn)
	carAvailability, err := client.GetAvailability(context.Background(), &availabilitypb.AvailabiltyRequest{
		Source: &availabilitypb.Location{
			Latitude:  2,
			Longitude: 3,
		},
	})
	if err != nil {
		log.Fatalf("error occured while calling Get Availability of car:%v", err)
	}
	for {
		msg, err := carAvailability.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error occured while reading stream %v", err)
		}

		a := Car{CarType: msg.GetCarType(), Location: msg.GetLocation(), Distance: msg.GetDistance()}
		carpool = append(carpool, a)
		log.Printf("receive:%v\n", msg)
	}
	c.IndentedJSON(http.StatusOK, &carpool)
}

func fetchAvailableCarNear(c *gin.Context) {
	waitResponse := make(chan struct{})
	var carpool []Car
	sAddress := ":3000"
	conn, e := grpc.Dial(sAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()
	client := routingpb.NewRoutingServiceClient(conn)
	carAvailability, err := client.GetAvailability(context.Background())
	if err != nil {
		log.Fatalf("error occured while calling Get Availability of car:%v", err)
	}
	req := routingpb.AvailabiltyRequest{
		Source: &routingpb.Location{
			Latitude:  3,
			Longitude: 2,
		},
	}
	if err := carAvailability.Send(&req); err != nil {
		log.Fatalf("Error while send to GetAvailability Server: %v", err)
	}
	if err := carAvailability.CloseSend(); err != nil {
		log.Fatalf("Error while close send to Get Availability Server: %v", err)
	}
	var carpool []AvailableResponse
	go func() {
		for {
			msg, err := carAvailability.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error occured while reading stream %v", err)
			}

			resRouting := routingpb.ListAvailabiltyResponse{
				AvailabiltyResponse: msg.GetAvailabiltyResponse(),
			}
			carpool = append(carpool, resRouting)
			log.Printf("receive:%v\n", msg)
		}
		c.IndentedJSON(http.StatusOK, &carpool)
		close(waitResponse)
	}()
}
