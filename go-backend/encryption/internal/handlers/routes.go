package handlers

import (
	"github.com/cryptofortress/backend/encryption/internal/services"
	"github.com/cryptofortress/backend/encryption/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the encryption service
func RegisterRoutes(router *gin.Engine, services *services.Services) {
	// Create handlers
	encryptionHandler := NewEncryptionHandler(services.Encryption)
	fpeHandler := NewFPEHandler(services.FPE)

	// Public routes (no authentication required for demo purposes)
	public := router.Group("/api/v1/encryption")
	{
		public.POST("/encrypt", encryptionHandler.Encrypt)
		public.POST("/decrypt", encryptionHandler.Decrypt)
		public.POST("/generate-key", encryptionHandler.GenerateKey)
		
		// FPE routes
		public.POST("/fpe/encrypt", fpeHandler.FPEEncrypt)
		public.POST("/fpe/decrypt", fpeHandler.FPEDecrypt)
		
		// Credit card routes
		public.POST("/credit-card/encrypt", fpeHandler.EncryptCreditCard)
		public.POST("/credit-card/decrypt", fpeHandler.DecryptCreditCard)
		
		// SSN routes
		public.POST("/ssn/encrypt", fpeHandler.EncryptSSN)
		public.POST("/ssn/decrypt", fpeHandler.DecryptSSN)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Encryption Service is running",
		})
	})
}