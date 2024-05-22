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
	log.Print("New try of login detected")
	// TODO expAt use config instead
	expAt := timestamppb.New(time.Now().Add(time.Hour))
	// TODO get user_id from db
	accessToken, err := jwt.CreateToken("5")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("New login token: ", accessToken)
	return &api.LoginResponse{AccessToken: accessToken,
		ExpiresAt: expAt}, nil
}

func (s *Server) ValidateToken(context.Context, *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	log.Print("New try of token validation detected")
	// TODO use ValidateToken func instead
	isValid := true
	// TODO take userID from token
	userID := "5"
	log.Print("Token validated: ", userID)
	return &api.ValidateTokenResponse{IsValid: isValid,
		UserId: userID}, nil
}

func (s *Server) Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error) {
	log.Print("New try of register detected")
	// TODO email, username and password validation
	// TODO get new user id from db
	userID := "5"
	createdAt := timestamppb.New(time.Now())

	log.Print("New user registered ID: ", userID)
	return &api.RegisterResponse{UserId: userID,
		CreatedAt: createdAt}, nil
}
