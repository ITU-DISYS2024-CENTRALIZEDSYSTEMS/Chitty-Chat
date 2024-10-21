package main

import (
 "context"
 pb "github.com/jgjo/grpc-go/entry"
 "google.golang.org/grpc"
 "google.golang.org/grpc/credentials/insecure"
 "log"
 "time"
)

func main() {
 conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
 if err != nil {
  log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
 }
 defer conn.Close()
 c := pb.NewBroadcastServiceClient(conn)

 ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 defer cancel()
 

 r, err := c.Broadcast(ctx, &pb.BroadcastRequest{Message: "Hello from gRPC client!"})
 if err != nil {
  log.Fatalf("error calling function SayHello: %v", err)
 }

 log.Printf("Response from gRPC server's SayHello function: %s", r.GetMessage())
}