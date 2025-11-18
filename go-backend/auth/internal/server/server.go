package server

import (
	"context"
	"net/http"
	"time"

	"github.com/cryptofortress/backend/auth/internal/config"
	"github.com/cryptofortress/backend/auth/internal/handlers"
	"github.com/cryptofortress/backend/auth/internal/middleware"
	"github.com/cryptofortress/backend/auth/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Server represents the authentication service server
type Server struct {
	config   *config.Config
	router   *gin.Engine
	services *services.Services
}

// New creates a new authentication server instance
func New(cfg *config.Config) *Server {
	// Initialize services
	authService := services.NewAuthService(cfg)
	mfaService := services.NewMFAService(cfg)
	rbacService := services.NewRBACService(cfg)
	
	services := &services.Services{
		Auth: authService,
		MFA:  mfaService,
		RBAC: rbacService,
	}
	
	// Create router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logging())
	
	// Register routes
	handlers.RegisterRoutes(router, services)
	
	return &Server{
		config:   cfg,
		router:   router,
		services: services,
	}
}

// Start begins serving requests
func (s *Server) Start() error {
	server := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.router,
	}
	
	// Graceful shutdown
	go func() {
		<-context.Background().Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		if err := server.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Server shutdown failed")
		}
	}()
	
	return server.ListenAndServe()
}

// Stop gracefully shuts down the server
func (s *Server) Stop() error {
	// Cleanup resources if needed
	return nil
}