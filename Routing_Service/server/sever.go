package server

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/Dinesh789kumar12/POC1/CarAppPb/routingpb"
	"google.golang.org/grpc"
)

type server struct {
	routingpb.UnimplementedRoutingServiceServer
}

func GetRates(req routingpb.RoutingService_GetRatesServer) error {
	return nil
}

func Booking(routingpb.RoutingService_BookingServer) error {
	return nil
}

func Confirmation(context.Context, *routingpb.ConfirmRequest) (*routingpb.ConfirmResponse, error) {
	return nil, nil
}
func (*server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
	log.Println("GetAvailability function called")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		log.Println(req.GetSource())
		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}

		// Build and send response to the client
		res := stream.Send(&routingpb.AvailabiltyResponse{
			CarType:  "audi",
			Distance: 6,
		})

		if res != nil {
			log.Fatalf("Error when response was sent to the client: %v", res)
		}
	}
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Error occured while registering address %v", err)
	}
	log.Println("Server Start Listening on Port :0.0.0.0:3000")
	s := grpc.NewServer()
	routingpb.RegisterRoutingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
