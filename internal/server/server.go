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
	// TODO expAt use config instead
	expAt := timestamppb.New(time.Now().Add(time.Hour))
	// TODO get user_id from db
	accessToken, err := jwt.CreateToken("5")
	if err != nil {
		log.Fatal(err)
	}
	return &api.LoginResponse{AccessToken: accessToken,
		ExpiresAt: expAt}, nil
}

func (s *Server) ValidateToken(context.Context, *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	// TODO use ValidateToken func instead
	isValid := true
	// TODO take userID from token
	userID := "5"
	return &api.ValidateTokenResponse{IsValid: isValid,
		UserId: userID}, nil
}

func (s *Server) Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error) {
	// TODO get new user id from db
	userID := "5"
	createdAt := timestamppb.New(time.Now())
	return &api.RegisterResponse{UserId: userID,
		CreatedAt: createdAt}, nil
}
