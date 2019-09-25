package main

import (
	"log"
	"net/http"

	pb "github.com/abvarun226/notes-app/note-svc/proto"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50051"
	address       = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	h := Handler{NoteService: pb.NewNoteClient(conn)}

	r := chi.NewRouter()
	r.Get("/note/v1/notes/{user}", h.ListNotesByUser)

	log.Printf("starting http server on port %s", address)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal("served failed to start")
	}
}
