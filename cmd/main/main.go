package main

import (
	"context"
	"github.com/437d5/jwt-auth/internal/app"
	"github.com/437d5/jwt-auth/internal/config"
	"github.com/joho/godotenv"
	"log"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Cannot initialize config: %s", err)
	}

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalf("Error creating new app: %s", err)
	}

	if err = a.Run(ctx); err != nil {
		log.Fatalf("Error starting new app: %s", err)
	}
}
