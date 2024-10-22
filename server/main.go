package main

import (
	"io"
	"log"
	"net"
	"sync"
	pb "github.com/jgjo/grpc-go/entry"
	"google.golang.org/grpc"
)

type server struct {
 pb.UnimplementedBroadcastServiceServer
 mu sync.Mutex
 clients map[string]pb.BroadcastService_BroadcastServer
 time int32
}



func (s *server) Broadcast(srv pb.BroadcastService_BroadcastServer) error {
    ctx := srv.Context()
    for{
        
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

        	// receive data from stream
		req, err := srv.Recv()
        if err == io.EOF {
            // return will close stream from server side
            log.Println("exit")
            return nil
        }
        if err != nil {
            log.Printf("receive error %v", err)
            continue
        }

        // If new user, add to clients
        if s.clients[req.GetUser()] == nil {
            s.mu.Lock()
            s.clients[req.GetUser()] = srv
            log.Printf("User %s has joined", req.User)
            s.mu.Unlock()
            continue
        }

        // if user leaves remove from clients
        defer func() {
            s.mu.Lock()
                delete(s.clients, req.GetUser())
                log.Printf("User %v has left", req.GetUser())
            s.mu.Unlock()
        }()

        // 
        s.mu.Lock()
        if(req.Timestamp > s.time){
            s.time = req.Timestamp + 1
        }else {
            s.time++
        }
        s.mu.Unlock()

        // send message to all clients
        log.Printf("Broadcasting message from | %v at time: %v", req.GetUser(), s.time)
        for user, client := range s.clients {
            if user != req.GetUser() {
                if err := client.Send(&pb.BroadcastMessage{
                    User: req.User,
                    Message: req.Message,
                    Timestamp: s.time,
                }); err != nil {
                    log.Printf("send error %v", err)
                }
            }
        }

    }
}



func main() {
    server := &server{
        clients: make(map[string]pb.BroadcastService_BroadcastServer),
        time: 0,
    }
    // create listener
lis, err := net.Listen("tcp", ":50005")
if err != nil {
    log.Fatalf("failed to listen: %v", err)
}

// create grpc server
s := grpc.NewServer()
    pb.RegisterBroadcastServiceServer(s, server)
    log.Println("server is running on port :50005")
// and start...
if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
}
}