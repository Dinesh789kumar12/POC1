package server

import (
	"context"
	"io"
	"log"
	"net"
	"sync"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/availabilitypb"
	"github.com/Dinesh789kumar12/POC1/CarAppPb/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	routingpb.UnimplementedRoutingServiceServer
}

type routApp struct {
	routeServer        grpc.Server
	routeListner       net.Listener
	availabilityClient availabilitypb.AvailabiltyServiceClient
}

var RoutApp routApp

func main() {
	RoutApp.init_app()

	if err := RoutApp.routeServer.Serve(RoutApp.routeListner); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *routApp) init_app() {
	var err error
	a.routeListner, err = net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Started listening on port 3000")
	routeServer := grpc.NewServer()
	routingpb.RegisterRoutingServiceServer(routeServer, &server{})

	log.Println("Client dialing on port 50052 for Availability ms")
	con_availability, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	a.availabilityClient = availabilitypb.NewAvailabiltyServiceClient(con_availability)

}

func (*server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	c := make(chan *routingpb.AvailabiltyRequest, 5)
	q := make(chan bool)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err != nil {
				q <- true
				break
			} else {
				c <- req
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c:
				v := <-c
				res, err := RoutApp.availabilityClient.GetAvailability(context.Background(), &availabilitypb.AvailabiltyRequest{
					Source: &availabilitypb.Location{
						Latitude:  v.GetSource().GetLatitude(),
						Longitude: v.Source.GetLongitude(),
					},
				})
				if err != nil {
					log.Fatalf("Unable to connnect Availability Service:%v", err)
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
					resRouting := routingpb.AvailabiltyResponse{
						CarType:  msg.GetCarType(),
						Location: msg.GetLocation(),
						Distance: msg.GetDistance(),
					}
					if err := stream.Send(&resRouting); err != nil {
						log.Fatalf("Error while send to GreetAll RPC: %v", err)
					}
				}
				log.Println("Data sending to Client.....")
			case <-q:
				log.Println("Exit code ............")
			}
		}
	}()
	wg.Wait()
	return nil
}

// func GetRates(req routingpb.RoutingService_GetRatesServer) error {
// 	return nil
// }

// func Booking(routingpb.RoutingService_BookingServer) error {
// 	return nil
// }

// func Confirmation(context.Context, *routingpb.ConfirmRequest) (*routingpb.ConfirmResponse, error) {
// 	return nil, nil
// }
// func (*server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
// 	log.Println("GetAvailability function called")
// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		log.Println(req.GetSource())
// 		if err != nil {
// 			log.Fatalf("Error when reading client request stream: %v", err)
// 		}

// 		// Build and send response to the client
// 		res := stream.Send(&routingpb.AvailabiltyResponse{
// 			CarType:  ,
// 			Distance: ,
// 		})

// 		if res != nil {
// 			log.Fatalf("Error when response was sent to the client: %v", res)
// 		}
// 	}
// }
