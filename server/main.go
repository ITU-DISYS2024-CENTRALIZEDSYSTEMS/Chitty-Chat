package main

import (
	chitty_chat "chitty-chat/chitty-chat"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type client struct {
	name string
	connection chitty_chat.ChittyChat_JoinConversationServer
}

type chittyChatServer struct {
	chitty_chat.UnimplementedChittyChatServer
	clients map[string]client
	lamport int32
	mu sync.Mutex
}

func (s *chittyChatServer) JoinConversation(stream chitty_chat.ChittyChat_JoinConversationServer) error {
	var author struct {
		id string ""
		name string ""
	}

	for {
		incomingMessage, err := stream.Recv()
		if err != nil {
			s.lamport++
			log.Println("Info |", author.name, "| Dropped the connection! | Lamport time", s.lamport)
			break
		}

		if (incomingMessage.Timestamp > s.lamport) {
			s.lamport = incomingMessage.Timestamp + 1
		} else {
			s.lamport++
		}

		if author.id == "" {
			author.id = incomingMessage.Author
			author.name = incomingMessage.Content
			s.mu.Lock()
			s.clients[author.id] = client{
				name: author.name,
				connection: stream,
			}
			s.mu.Unlock()

			log.Println("Info |", author.name, "| Joined Chitty-Chat | Lamport time", s.lamport)
			s.broadcastMessage(&chitty_chat.Message{
				Author: s.clients[author.id].name,
				Content: "joined Chitty-Chat",
				Timestamp: s.lamport,
			})
			continue
		}

		if len(incomingMessage.Content) <= 128 {
			incomingMessage.Author = s.clients[incomingMessage.Author].name
			incomingMessage.Timestamp = s.lamport
			log.Println("Info |", author.name, "| Sent a message | Lamport time", s.lamport)
			s.broadcastMessage(incomingMessage)
		} else {
			log.Println("Info |", author.name, "| Tried to send a too big message!")
		}
	}

	s.lamport++

	s.mu.Lock()
	delete(s.clients, author.id)
	s.mu.Unlock()

	s.broadcastMessage(&chitty_chat.Message{
		Author: author.name,
		Content: "left Chitty-Chat",
		Timestamp: s.lamport,
	})

	return nil
}

func (s *chittyChatServer) broadcastMessage(msg *chitty_chat.Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lamport++
	msg.Timestamp = s.lamport
	log.Println("Broadcasting |", msg.Author, "|", msg.Content, "| Lamport time", msg.Timestamp)

	for _, client := range s.clients {
		if err := client.connection.Send(msg); err != nil {
			log.Println("Error sending message to a client:", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Cannot create listener: ", err)
	}

	server := grpc.NewServer()
	service := &chittyChatServer{
		lamport: 0,
		clients: make(map[string]client),
	}

	chitty_chat.RegisterChittyChatServer(server, service)
	
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("Cannot serve service:", err)
	}
}