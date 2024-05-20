package main

import (
	"github.com/437d5/jwt-auth/internal/server"
	"github.com/437d5/jwt-auth/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.Server{}
	api.RegisterAuthServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
