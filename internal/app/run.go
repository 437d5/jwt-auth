package app

import (
	"context"
	"fmt"
	"github.com/437d5/jwt-auth/internal/config"
	"github.com/437d5/jwt-auth/internal/db"
	"github.com/437d5/jwt-auth/internal/server"
	"github.com/437d5/jwt-auth/pkg/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type App struct {
	DB       *mongo.Client
	SRV      *grpc.Server
	Listener net.Listener
}

func databaseConnect(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	connString := cfg.CreateConnString()

	clientOptions := options.Client().ApplyURI(connString)

	log.Print("Trying to connect to database")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {

		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Print("Connected to database")
	return client, nil
}

func (a *App) Run(ctx context.Context) error {
	ch := make(chan error, 1)

	go func() {
		log.Print("Starting new server")
		err := a.SRV.Serve(a.Listener)
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		go func() {
			a.SRV.GracefulStop()
			ch <- nil
		}()

		select {
		case err := <-ch:
			return err
		case <-shutdownCtx.Done():
			return fmt.Errorf("shutdown timed out")
		}
	}
}

func NewServer(cfg *config.Config, client *mongo.Client) (*grpc.Server, net.Listener, error) {
	s := grpc.NewServer()
	srv := &server.Server{
		DB: db.Repo{
			DB: client,
		},
	}
	api.RegisterAuthServiceServer(s, srv)
	portStr := fmt.Sprintf(":%s", cfg.SRV.Port)

	l, err := net.Listen("tcp", portStr)
	if err != nil {
		return nil, nil, err
	}
	return s, l, nil
}

func NewApp(ctx context.Context, cfg *config.Config) (*App, error) {
	client, err := databaseConnect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	s, l, err := NewServer(cfg, client)
	if err != nil {
		return nil, err
	}

	a := &App{
		DB:       client,
		SRV:      s,
		Listener: l,
	}
	return a, nil
}
