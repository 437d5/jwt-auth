package server

import (
	"context"
	"errors"
	"github.com/437d5/jwt-auth/internal/db"
	"github.com/437d5/jwt-auth/internal/jwt"
	"github.com/437d5/jwt-auth/internal/validations"
	"github.com/437d5/jwt-auth/pkg/api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type Server struct {
	api.UnimplementedAuthServiceServer
	DB db.Repo
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

func (s *Server) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	log.Print("New try of register detected")
	user := db.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	ok := validations.ValidateUsername(user.Username)
	if !ok {
		return nil, errors.New("invalid username")
	}
	ok = validations.ValidatePassword(user.Password)
	if !ok {
		return nil, errors.New("invalid password")
	}
	ok = validations.ValidateEmail(user.Email)
	if !ok {
		return nil, errors.New("invalid email")
	}
	err := s.DB.CreateNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	// TODO get new user id from db
	userID := "5"
	createdAt := timestamppb.New(time.Now())

	log.Print("New user registered ID: ", userID)
	return &api.RegisterResponse{UserId: userID,
		CreatedAt: createdAt}, nil
}
