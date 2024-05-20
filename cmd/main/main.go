package main

import (
	"context"
	"github.com/437d5/jwt-auth/internal/app"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := app.RunServer(ctx)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
