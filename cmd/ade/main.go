package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kig3-svg/agentic-ai-ade/internal/app"
	"github.com/kig3-svg/agentic-ai-ade/internal/config"
	"github.com/kig3-svg/agentic-ai-ade/pkg/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ade",
	Short: "Agentic AI Development Environment",
	Long:  "A Go-based platform for autonomous AI agents",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServer(cmd.Context())
	},
}

func runServer(ctx context.Context) error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Initialize logger
	log := logger.NewLogger(cfg.LogLevel)
	log.Info("Starting Agentic AI ADE...")
	log.Infof("Environment: %s", cfg.Environment)

	// Initialize application
	application, err := app.NewApp(ctx, cfg, log)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
		return err
	}

	// Start application
	if err := application.Start(ctx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
		return err
	}

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info("Shutting down gracefully...")

	if err := application.Stop(ctx); err != nil {
		log.Errorf("Error during shutdown: %v", err)
		return err
	}

	log.Info("Agentic AI ADE stopped successfully")
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
