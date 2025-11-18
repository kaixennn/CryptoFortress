package main

import (
	"log"
	"os"

	"github.com/cryptofortress/backend/encryption/internal/server"
	"github.com/cryptofortress/backend/encryption/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize server
	srv := server.New(cfg)

	// Start server
	log.Printf("Starting Encryption Service on port %s", cfg.Port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}