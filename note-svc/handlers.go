package main

import (
	"context"
	"log"

	pb "github.com/abvarun226/notes-app/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler is used to implement handlers of note server.
type Handler struct{}

// ListNotes implements note.NoteServer
func (s *Handler) ListNotes(ctx context.Context, in *pb.NotesRequest) (*pb.NotesResponse, error) {
	notes := make([]string, 0)

	switch in.UserId {
	case "varun":
		notes = append(notes, "tutorial on grpc", "tutorial on istio", "tutorial on k8s")
	case "ayada":
		notes = append(notes, "tutorial on python", "tutorial on golang")
	default:
		log.Printf("user not found: %s", in.UserId)
		return nil, status.New(codes.NotFound, "user not found").Err()
	}

	return &pb.NotesResponse{Notes: notes}, nil
}
