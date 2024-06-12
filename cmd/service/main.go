package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/shankssc/Stripe-microservice/internal/application"
	"github.com/shankssc/Stripe-microservice/internal/config"
	"github.com/shankssc/Stripe-microservice/pkg/logger"
)

func main() {
	// Initialize configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger := logger.NewLogger(cfg.LogLevel)

	// Create a new application
	app := application.NewApp(cfg, logger)

	// Create context that listens for interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Start the application
	if err := app.Run(ctx); err != nil {
		log.Fatalf("Failed to run the application: %v", err)
	}

	logger.Info("Application stopped")
}
