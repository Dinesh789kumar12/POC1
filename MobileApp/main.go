package main

import (
	"context"
	"log"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/availabilitypb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	router := gin.Default()
	rg := router.Group("api/v1/carapp")
	{
		rg.GET("/:id", fetchAvailableCarNearby)
	}

	router.Run()
}

func fetchAvailableCarNearby(c *gin.Context) {
	sAddress := ":50052"
	conn, e := grpc.Dial(sAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := availabilitypb.NewAvailabiltyServiceClient(conn)
	carAvailability, e := client.GetAvailability(context.Background(), &availabilitypb.AvailabiltyRequest{
		Source: &availabilitypb.Location{
			Latitude:  2,
			Longitude: 3,
		},
	})
	if e != nil {
		log.Fatalf("Failed to get car: %v", e)
	}

	c.JSON(200, &carAvailability)
}
