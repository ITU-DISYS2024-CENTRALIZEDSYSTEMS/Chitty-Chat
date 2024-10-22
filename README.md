# How to use

Chitty Chat is split into two parts, the server and the client. The server is responsible for a single chat where multiple clients can connect to. Clients can connect and disconnect as they please, but when they are connected the server will broadcast messages to all connected clients.

## Server

Run the server using the following command:

```go
go run ./server/main.go
```

## Client

Run a client using the following command:

```go
go run ./client/main.go
```

Then you type the id you want to use for session followed by a display name.

The standard input stream is then used to send messages to the server.
