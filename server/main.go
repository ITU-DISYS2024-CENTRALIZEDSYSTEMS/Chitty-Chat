package main

import (
	chitty_chat "chitty-chat/chitty-chat"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type chittyChatServer struct {
	chitty_chat.UnimplementedChittyChatServer
}

func (s chittyChatServer) PublishMessage(context.Context, *chitty_chat.TextMessage) (*chitty_chat.TextMessage, error) {
	return nil, nil
}

func (s chittyChatServer) JoinConversation(context.Context, *chitty_chat.Message) (*chitty_chat.Message, error) {
	return nil, nil
}

func (s chittyChatServer) LeaveConversation(context.Context, *chitty_chat.Message) (*chitty_chat.Message, error) {
	return nil, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	server := grpc.NewServer()
	service := &chittyChatServer{}

	chitty_chat.RegisterChittyChatServer(server, service)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Cannot serve service: %s", err)
	}
}