package handlers

import (
	"github.com/cryptofortress/backend/keymgmt/internal/services"
	"github.com/cryptofortress/backend/keymgmt/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the key management service
func RegisterRoutes(router *gin.Engine, services *services.Services) {
	// Create handlers
	keyHandler := NewKeyHandler(services.Key)
	rotationHandler := NewRotationHandler(services.Rotation)
	shamirHandler := NewShamirHandler(services.Shamir)
	replicationHandler := NewReplicationHandler(services.Replication)

	// Public routes (no authentication required for demo purposes)
	public := router.Group("/api/v1/keymgmt")
	{
		// Key management routes
		public.POST("/keys/generate", keyHandler.GenerateKey)
		public.POST("/keys/generate-pair", keyHandler.GenerateKeyPair)
		public.POST("/keys/store", keyHandler.StoreKey)
		public.POST("/keys/retrieve", keyHandler.RetrieveKey)
		public.POST("/keys/delete", keyHandler.DeleteKey)
		
		// Key rotation routes
		public.POST("/rotation/rotate", rotationHandler.RotateKey)
		public.POST("/rotation/schedule", rotationHandler.ScheduleRotation)
		public.POST("/rotation/cancel", rotationHandler.CancelRotation)
		public.POST("/rotation/schedule/get", rotationHandler.GetRotationSchedule)
		public.POST("/rotation/auto-enable", rotationHandler.EnableAutoRotation)
		public.POST("/rotation/auto-disable", rotationHandler.DisableAutoRotation)
		
		// Shamir's Secret Sharing routes
		public.POST("/shamir/split", shamirHandler.SplitSecret)
		public.POST("/shamir/combine", shamirHandler.CombineShares)
		public.POST("/shamir/distribute", shamirHandler.DistributeKey)
		public.POST("/shamir/recover", shamirHandler.RecoverKey)
		
		// Key replication routes
		public.POST("/replication/replicate", replicationHandler.ReplicateKey)
		public.POST("/replication/enable", replicationHandler.EnableCrossRegionReplication)
		public.POST("/replication/disable", replicationHandler.DisableCrossRegionReplication)
		public.POST("/replication/backup", replicationHandler.BackupKey)
		public.POST("/replication/restore", replicationHandler.RestoreKey)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Key Management Service is running",
		})
	})
}