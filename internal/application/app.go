package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shankssc/Stripe-microservice/internal/config"
	"github.com/shankssc/Stripe-microservice/pkg/logger"
)

type App struct {
	Config *config.Config
	Logger *logger.Logger
	Router http.Handler
	Server *http.Server
	rdb    *redis.Client
}

func NewApp(cfg *config.Config, logger *logger.Logger) *App {
	router := setupRouter()

	server := &http.Server{
		Addr:    "localhost:" + cfg.ServerPort,
		Handler: router,
	}
	return &App{
		Config: cfg,
		Logger: logger,
		Router: router,
		Server: server,
		rdb: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (a *App) Run(ctx context.Context) error {

	err := a.rdb.Ping(ctx).Err()

	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	a.Logger.Info("Starting server on port", a.Config.ServerPort)

	ch := make(chan error, 1)

	// Run server in a separate goroutine to enable graceful shutdown
	go func() {
		if err = a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ch <- fmt.Errorf("server shutdown error: %v", err)
		}
		close(ch)
	}()

	// var err error

	select {
	case err = <-ch:
		// Logging the error if it occurred during server startup
		a.Logger.Error("server startup error:", err)
		return err
	// Wait for context cancellation
	case <-ctx.Done():
		// Attempt to gracefully shut down the server
		a.Logger.Info("Shutting down server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Try to gracefully shut down the server
		if shutdownErr := a.Server.Shutdown(shutdownCtx); shutdownErr != nil {
			a.Logger.Error("server shutdown error:", shutdownErr)
			return shutdownErr
		}
	}
	return nil
}
