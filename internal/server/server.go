package server

import (
	"context"
	"errors"
	"github.com/437d5/jwt-auth/internal/db"
	"github.com/437d5/jwt-auth/internal/jwt"
	"github.com/437d5/jwt-auth/internal/pswd"
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
	log.Printf("Username: %s, password: %s", req.GetUsername(), req.GetPassword())
	// TODO expAt use config instead
	expAt := timestamppb.New(time.Now().Add(time.Hour))
	user, err := s.DB.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		log.Print(err)
		return nil, err
	}
	ok := pswd.Compare(req.GetPassword(), user.Password)
	if !ok {
		log.Print("Password incorrect")
		return nil, errors.New("wrong password")
	}
	accessToken, err := jwt.CreateToken(user.ID)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	log.Printf("New login with id: %s", user.ID.Hex())
	return &api.LoginResponse{AccessToken: accessToken,
		ExpiresAt: expAt}, nil
}

func (s *Server) ValidateToken(ctx context.Context, req *api.ValidateTokenRequest) (*api.ValidateTokenResponse, error) {
	log.Print("New try of token validation detected")
	// TODO change secret to config
	isValid := false
	token, err := jwt.ValidToken(req.AccessToken, []byte("secret"))
	if err != nil {
		log.Print(err)
		return nil, err
	} else {
		isValid = true
	}

	userID, err := jwt.GetIDFromToken(token)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	log.Print("Token validated: ", userID)
	return &api.ValidateTokenResponse{IsValid: isValid,
		UserId: userID.Hex()}, nil
}

func (s *Server) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	log.Print("New try of register detected")
	user := db.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.Email,
	}
	ok := validations.ValidateUsername(user.Username)
	if !ok {
		log.Print("Username incorrect")
		return nil, errors.New("invalid username")
	}
	ok = validations.ValidatePassword(user.Password)
	if !ok {
		log.Print("Password incorrect")
		return nil, errors.New("invalid password")
	}
	ok = validations.ValidateEmail(user.Email)
	if !ok {
		log.Print("Email incorrect")
		return nil, errors.New("invalid email")
	}
	err := s.DB.CreateNewUser(ctx, user)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// TODO get new user id from db
	newUser, err := s.DB.GetUserByUsername(ctx, user.Username)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	userID := newUser.ID.Hex()
	createdAt := timestamppb.New(time.Now())

	log.Print("New user registered ID: ", userID)
	return &api.RegisterResponse{UserId: userID,
		CreatedAt: createdAt}, nil
}
