package main

import (
	"log"
	"net"

	pb "github.com/abvarun226/note-svc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	srv, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	h := Handler{}

	s := grpc.NewServer()
	pb.RegisterNoteServer(s, &h)

	log.Printf("starting note server on port %s", port)
	if err := s.Serve(srv); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
