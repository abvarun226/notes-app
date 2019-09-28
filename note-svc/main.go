package main

import (
	"log"
	"net"
	"strings"

	pb "github.com/abvarun226/notes-app/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Server Address.
const (
	ServerAddress = ":50051"
)

func main() {
	viper.SetConfigName("notes")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("server.address", ServerAddress)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	serverAddress := viper.GetString("server.address")

	srv, err := net.Listen("tcp", ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	h := Handler{}

	s := grpc.NewServer()
	pb.RegisterNoteServer(s, &h)

	log.Printf("starting note server on port %s", serverAddress)
	if err := s.Serve(srv); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
