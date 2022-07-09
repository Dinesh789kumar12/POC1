package main

import (
	"log"
	"net"
	"time"

	data "github.com/Dinesh789kumar12/POC1/Availabilty_Service/data"
	service "github.com/Dinesh789kumar12/POC1/Availabilty_Service/service"
	"github.com/Dinesh789kumar12/POC1/CarAppPb/availabilitypb"
	"google.golang.org/grpc"
)

type server struct {
	availabilitypb.UnimplementedAvailabiltyServiceServer
}

func (*server) GetAvailability(req *availabilitypb.AvailabiltyRequest, stream availabilitypb.AvailabiltyService_GetAvailabilityServer) error {
	for {
		source := req.GetSource()
		x := source.GetLatitude()
		y := source.GetLongitude()
		userlocation := data.LocationName[x][y]
		for _, c := range data.Carpool {
			var a, b int
			if c.Available {
				a, b = service.RandomNumberGenerator()
				CarLocation := data.LocationName[a][b]
				m := service.Distance(a, b, userlocation)
				res := &availabilitypb.AvailabiltyResponse{
					CarType:  c.Model,
					Location: CarLocation,
					Distance: m,
				}
				stream.Send(res)
				log.Printf("Sent:%v", res)
				time.Sleep(20 * time.Nanosecond)
			}
		}
		time.Sleep(20 * time.Second)
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Error occured while registering address %v", err)
	}
	log.Println("Server Start Listening on Port :0.0.0.0:50052")
	register := grpc.NewServer()
	availabilitypb.RegisterAvailabiltyServiceServer(register, &server{})

	if err := register.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
