package server

import (
	"context"
	"net/http"
	"time"

	"github.com/cryptofortress/backend/encryption/internal/config"
	"github.com/cryptofortress/backend/encryption/internal/handlers"
	"github.com/cryptofortress/backend/encryption/internal/middleware"
	"github.com/cryptofortress/backend/encryption/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Server represents the encryption service server
type Server struct {
	config   *config.Config
	router   *gin.Engine
	services *services.Services
}

// New creates a new encryption server instance
func New(cfg *config.Config) *Server {
	// Initialize services
	encryptionService := services.NewEncryptionService(cfg)
	fpeService := services.NewFPEService(cfg)
	
	services := &services.Services{
		Encryption: encryptionService,
		FPE:        fpeService,
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