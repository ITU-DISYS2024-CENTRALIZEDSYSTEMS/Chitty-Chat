package main

import (
 "context"
 pb "github.com/jgjo/grpc-go/entry"
 "google.golang.org/grpc"
 "log"
 "net"
)

type server struct {
 pb.UnimplementedBroadcastServiceServer
}


func (s *server) Broadcast(ctx context.Context, in *pb.BroadcastRequest) (*pb.BroadcastResponse, error) {
	return &pb.BroadcastResponse{Message: in.Message}, nil
}

func main() {
 lis, err := net.Listen("tcp", ":50051")
 if err != nil {
  log.Fatalf("failed to listen on port 50051: %v", err)
 }

 s := grpc.NewServer()
 pb.RegisterBroadcastServiceServer(s, &server{})
 log.Printf("gRPC server listening at %v", lis.Addr())

 if err := s.Serve(lis); err != nil {
  log.Fatalf("failed to serve: %v", err)
 }
}