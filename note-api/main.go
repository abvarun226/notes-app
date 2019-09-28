package main

import (
	"log"
	"net/http"
	"strings"

	pb "github.com/abvarun226/notes-app/proto"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Server Address.
const (
	NoteSVCAddress = "note-svc:50051"
	ServerAddress  = ":50052"
)

func main() {
	viper.SetConfigName("notes")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("notesvc.address", NoteSVCAddress)
	viper.SetDefault("server.address", ServerAddress)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	svcAddress := viper.GetString("notesvc.address")
	serverAddress := viper.GetString("server.address")

	// Set up a connection to the server.
	conn, err := grpc.Dial(svcAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	h := Handler{NoteService: pb.NewNoteClient(conn)}

	r := chi.NewRouter()
	r.Get("/note/v1/notes/{user}", h.ListNotesByUser)

	log.Printf("starting http server on port %s", serverAddress)
	if err := http.ListenAndServe(serverAddress, r); err != nil {
		log.Fatal("served failed to start")
	}
}
