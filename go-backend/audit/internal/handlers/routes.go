package handlers

import (
	"github.com/cryptofortress/backend/audit/internal/services"
	"github.com/cryptofortress/backend/audit/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the audit & compliance service
func RegisterRoutes(router *gin.Engine, services *services.Services) {
	// Create handlers
	auditHandler := NewAuditHandler(services.Audit)
	complianceHandler := NewComplianceHandler(services.Compliance)
	siemHandler := NewSIEMHandler(services.SIEM)

	// Public routes (no authentication required for demo purposes)
	public := router.Group("/api/v1/audit")
	{
		// Audit routes
		public.POST("/events/log", auditHandler.LogEvent)
		public.POST("/events/trail", auditHandler.GetAuditTrail)
		public.POST("/events/immutable", auditHandler.CreateImmutableTrail)
		public.POST("/events/verify", auditHandler.VerifyTrailIntegrity)
		
		// Compliance routes
		public.POST("/compliance/report", complianceHandler.GenerateReport)
		public.POST("/compliance/status", complianceHandler.GetComplianceStatus)
		public.GET("/compliance/standards", complianceHandler.ListComplianceStandards)
		public.POST("/compliance/gdpr/request", complianceHandler.HandleDataSubjectRequest)
		public.POST("/compliance/gdpr/inventory", complianceHandler.GenerateDataInventory)
		
		// SIEM routes
		public.POST("/siem/event", siemHandler.SendEvent)
		public.GET("/siem/threats", siemHandler.MonitorThreats)
		public.POST("/siem/endpoint", siemHandler.ConfigureEndpoint)
		public.POST("/siem/alert-rule", siemHandler.CreateAlertRule)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Audit & Compliance Service is running",
		})
	})
}