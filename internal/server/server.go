package server

import (
	"context"
	"github.com/437d5/jwt-auth/internal/jwt"
	"github.com/437d5/jwt-auth/pkg/api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type Server struct {
	api.UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	expAt := timestamppb.New(time.Now().Add(time.Hour))
	accessToken, err := jwt.CreateToken("5")
	if err != nil {
		log.Fatal(err)
	}
	return &api.LoginResponse{AccessToken: accessToken,
		ExpiresAt: expAt}, nil
}

func (s *Server) ValidateToken(context.Context, *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	return &api.ValidateTokenResponse{}, nil
}

func (s *Server) Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error) {
	return &api.RegisterResponse{}, nil
}
