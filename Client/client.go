package main

import (
	proto "Chitty-Chat/Gen"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
);

func main(){

	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()));
	
	if err != nil {
		log.Fatalf("An error occured during client setup: %s", err);
	}
	
	ctx := context.Background();

	fmt.Print("Joined the conversation!");

	client := proto.NewChittyChatClient(conn);

	user := proto.User{
		Name: "Kevin",
		TimeStampCreation: "69",
	}
	
	stream, err := client.JoinConversation(ctx);
	
	if err != nil {
		log.Fatalf("Couldn't join the conversation: %s", err);
	}
	

	go sendMessage(stream, &user);
	go recieveMessage(stream);
	select {}
}


func sendMessage(stream grpc.BidiStreamingClient[proto.Message, proto.Message], user *proto.User){

	for {
		input := bufio.NewScanner(os.Stdin)
    	input.Scan()
		if input.Text() == "" {
			return;
		}
    	msg := proto.Message {
			User: user,
			Message: input.Text(),
			TimeStamp: "10",
		}
		stream.Send(&msg);
	}
}

func recieveMessage(stream grpc.BidiStreamingClient[proto.Message, proto.Message]){
	for {
		revieced, err := stream.Recv();
		
		if err != nil{
			log.Fatalf("An error occured while listening: %s", err);
		}
		log.Print(revieced);
	}
}


