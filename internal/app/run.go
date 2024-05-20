package app

import (
	"context"
	"fmt"
	"github.com/437d5/jwt-auth/internal/server"
	"github.com/437d5/jwt-auth/pkg/api"
	"google.golang.org/grpc"
	"net"
	"time"
)

func RunServer(ctx context.Context) error {
	s := grpc.NewServer()
	srv := &server.Server{}
	api.RegisterAuthServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	ch := make(chan error, 1)

	go func() {
		err := s.Serve(l)
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		go func() {
			s.GracefulStop()
			ch <- nil
		}()

		select {
		case err = <-ch:
			return err
		case <-shutdownCtx.Done():
			return fmt.Errorf("shutdown timed out")
		}
	}
}
