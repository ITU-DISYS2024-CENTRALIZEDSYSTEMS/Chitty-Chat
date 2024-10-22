package main

import (
	"bufio"
	chitty_chat "chitty-chat/chitty-chat"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lamport int32 = 0

func receiveMessage(stream grpc.BidiStreamingClient[chitty_chat.Message, chitty_chat.Message]) {
	for {
		incomingMessage, err := stream.Recv()
		if err != nil {
			log.Fatalln("Error receiving message:", err)
		}

		if (incomingMessage.Timestamp > lamport) {
			lamport = incomingMessage.Timestamp + 1
		} else {
			lamport++
		}

		log.Println(incomingMessage.Author, "|", incomingMessage.Content, "| Lamport time", lamport)
	}
}

func sendMessage(stream grpc.BidiStreamingClient[chitty_chat.Message, chitty_chat.Message], id string) {	
	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		lamport++
		err := stream.Send(&chitty_chat.Message{
			Author:    id,
			Content:   input.Text(),
			Timestamp: lamport,
		})
		if err != nil {
			log.Fatalln("Error sending message:", err)
		}
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Please input your id:")
	input.Scan()
	var id string = input.Text()
	
	fmt.Println("Please input your name:")
	input.Scan()
	var name string = input.Text()

	grpcOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(":8080", grpcOptions)
	if err != nil {
		log.Fatalf("Cannot create client: %s", err)
	}

	client := chitty_chat.NewChittyChatClient(conn)

	ctx := context.Background()
	stream, err := client.JoinConversation(ctx)
	if err != nil {
		log.Fatalf("Cannot join conversation: %s", err)
	}

	err = stream.Send(&chitty_chat.Message{
		Author: id,
		Content: name,
		Timestamp: 1,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %s", err)
	}

	go receiveMessage(stream)
	go sendMessage(stream, id)

	select {}
}