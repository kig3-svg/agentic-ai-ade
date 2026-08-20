package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kig3-svg/agentic-ai-ade/internal/config"
	"github.com/kig3-svg/agentic-ai-ade/pkg/logger"
)

// Server represents the HTTP server
type Server struct {
	config *config.Config
	logger *logger.Logger
	server *http.Server
	mux    *http.ServeMux
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config, log *logger.Logger) *Server {
	mux := http.NewServeMux()
	
	server := &Server{
		config: cfg,
		logger: log,
		mux:    mux,
	}

	// Register routes
	server.registerRoutes()

	server.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: server.mux,
	}

	return server
}

// registerRoutes registers HTTP routes
func (s *Server) registerRoutes() {
	// Health check endpoint
	s.mux.HandleFunc("/health", s.handleHealth)

	// API endpoints
	s.mux.HandleFunc("/api/v1/agents", s.handleGetAgents)
	s.mux.HandleFunc("/api/v1/tasks", s.handleCreateTask)
}

// handleHealth handles health check requests
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy","version":"0.1.0"}`)
}

// handleGetAgents handles get agents requests
func (s *Server) handleGetAgents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"agents":[]}`)
}

// handleCreateTask handles create task requests
func (s *Server) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"taskId":"task-001","status":"pending"}`)
}

// Start starts the HTTP server
func (s *Server) Start(ctx context.Context) error {
	return s.server.ListenAndServe()
}

// Stop stops the HTTP server gracefully
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
