// package client

// import (
// 	"context"
// 	"fmt"
// 	"io"
// 	"log"
// 	"time"

// 	"github.com/Dinesh789kumar12/POC1/CarAppPb/routingpb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {
// 	con, err := grpc.Dial("0.0.0.0:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Error connecting: %v \n", err)
// 	}

// 	defer con.Close()
// 	c := routingpb.NewRoutingServiceClient(con)
// 	checkAvailability(c)

// }

// func checkAvailability(c routingpb.RoutingServiceClient) {

// 	// Get the stream and a possible error from the CreateUser function
// 	stream, err := c.GetAvailability(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error when getting stream object: %v", err)
// 		return
// 	}

// 	// Initialize the container struct and call the initUsers function
// 	// to get user objects to send on the request message.
// 	requests := container{}.initUsers()

// 	waitResponse := make(chan struct{})

// 	go func() {
// 		for _, req := range requests {
// 			stream.Send(req)
// 			time.Sleep(1000 * time.Millisecond)
// 		}
// 		stream.CloseSend()
// 	}()

// 	// Use a go routine to receive response messages from the server
// 	go func() {
// 		for {
// 			res, err := stream.Recv()
// 			if err == io.EOF {
// 				break
// 			}
// 			if err != nil {
// 				log.Fatalf("Error when receiving response: %v", err)
// 			}
// 			fmt.Println("Server Response: ", res)
// 		}
// 		close(waitResponse)
// 	}()
// 	<-waitResponse
// }

// type container struct {
// }

// func (c container) initUsers() {

// }
