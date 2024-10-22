package main

import (
    "bufio"
    "context"
    "log"
    "os"
    "strings"

    pb "github.com/jgjo/grpc-go/entry"
    "google.golang.org/grpc"
)

func main() {

	// dail server
	conn, err := grpc.NewClient(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
    defer conn.Close()
    c := pb.NewBroadcastServiceClient(conn)

    stream, err := c.Broadcast(context.Background())
    done := make(chan bool)
    if err != nil {
        log.Fatalf("error creating stream: %v", err)
    }

    // receive broadcast messages from server
    go func() {
       for{
         resp, err := stream.Recv()
            if err != nil {
                close(done)
                return
            }
            log.Printf("Message from %s | %s", resp.User ,resp.Message)

       }
    }()

    reader := bufio.NewReader(os.Stdin)

    log.Print("Enter username: ")
    user, err := reader.ReadString('\n')
    if err != nil {
        log.Fatalf("failed to read input: %v", err)
    }
    user = strings.TrimSpace(user)
    stream.Send(&pb.BroadcastMessage{
        User: user,
    })

    for {
        // Read input from the terminal
        log.Print("Enter message: ")
        message, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("failed to read input: %v", err)
        }
        message = strings.TrimSpace(message)

        stream.Send(&pb.BroadcastMessage{
            Message:   message,
            User:      user,
            Timestamp: 0,
        })
    }
}