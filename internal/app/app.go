package app

import (
	"context"

	"github.com/kig3-svg/agentic-ai-ade/internal/config"
	"github.com/kig3-svg/agentic-ai-ade/internal/agent"
	"github.com/kig3-svg/agentic-ai-ade/pkg/logger"
)

// App represents the main application
type App struct {
	config      *config.Config
	logger      *logger.Logger
	agentPool   *agent.Pool
	server      *Server
}

// NewApp creates a new application instance
func NewApp(ctx context.Context, cfg *config.Config, log *logger.Logger) (*App, error) {
	log.Info("Initializing application components...")

	// Create agent pool
	pool := agent.NewPool(cfg.Agent.MaxConcurrent, log)

	// Create server
	server := NewServer(cfg, log)

	app := &App{
		config:    cfg,
		logger:    log,
		agentPool: pool,
		server:    server,
	}

	return app, nil
}

// Start starts the application
func (a *App) Start(ctx context.Context) error {
	a.logger.Info("Starting application components...")

	// Start agent pool
	a.agentPool.Start(ctx)
	a.logger.Info("Agent pool started")

	// Start server
	go func() {
		if err := a.server.Start(ctx); err != nil {
			a.logger.Errorf("Server error: %v", err)
		}
	}()
	a.logger.Infof("Server started on %s:%d", a.config.Server.Host, a.config.Server.Port)

	return nil
}

// Stop stops the application gracefully
func (a *App) Stop(ctx context.Context) error {
	a.logger.Info("Stopping application components...")

	// Stop server
	if err := a.server.Stop(ctx); err != nil {
		a.logger.Errorf("Error stopping server: %v", err)
		return err
	}

	// Stop agent pool
	a.agentPool.Stop(ctx)
	a.logger.Info("Agent pool stopped")

	return nil
}

// GetAgentPool returns the agent pool
func (a *App) GetAgentPool() *agent.Pool {
	return a.agentPool
}

// GetLogger returns the logger
func (a *App) GetLogger() *logger.Logger {
	return a.logger
}

// GetConfig returns the configuration
func (a *App) GetConfig() *config.Config {
	return a.config
}
