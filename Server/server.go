package main

import (
	proto "Chitty-Chat/Gen"
	"log"
	"net"

	"google.golang.org/grpc"
);

type ChittyChatServer struct {
	proto.UnimplementedChittyChatServer
	
}

func (s *ChittyChatServer) JoinConversation(stream proto.ChittyChat_JoinConversationServer) error{
	for {
		revieced, err := stream.Recv();
		if err != nil{
			log.Println("An error occured while listening:", err);
		}
		log.Print(revieced.Message);
	}
}

func main(){

	grpcServer := grpc.NewServer();

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("An error occured during server listening: %s", err);
	}

	proto.RegisterChittyChatServer(grpcServer, &ChittyChatServer{});

	grpcServer.Serve(listener);

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error creating the server %v", err)
	}
}
