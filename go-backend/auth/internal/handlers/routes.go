package handlers

import (
	"github.com/cryptofortress/backend/auth/internal/services"
	"github.com/cryptofortress/backend/auth/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the authentication service
func RegisterRoutes(router *gin.Engine, services *services.Services) {
	// Create handlers
	authHandler := NewAuthHandler(services.Auth)
	mfaHandler := NewMFAHandler(services.MFA)
	rbacHandler := NewRBACHandler(services.RBAC)

	// Public routes (no authentication required)
	public := router.Group("/api/v1/auth")
	{
		public.POST("/login", authHandler.Login)
		public.POST("/refresh", authHandler.Refresh)
		public.POST("/register", authHandler.Register)
		public.POST("/logout", authHandler.Logout)
	}

	// Protected routes (authentication required)
	protected := router.Group("/api/v1/auth")
	protected.Use(middleware.AuthMiddleware(services.Auth))
	{
		// MFA routes
		mfa := protected.Group("/mfa")
		{
			mfa.POST("/totp/enable", mfaHandler.EnableTOTP)
			mfa.POST("/totp/verify", mfaHandler.VerifyTOTP)
			mfa.POST("/totp/disable", mfaHandler.DisableTOTP)
			
			mfa.POST("/webauthn/register", mfaHandler.RegisterWebAuthn)
			mfa.POST("/webauthn/register/verify", mfaHandler.VerifyWebAuthnRegistration)
			mfa.POST("/webauthn/authenticate", mfaHandler.AuthenticateWebAuthn)
			mfa.POST("/webauthn/authenticate/verify", mfaHandler.VerifyWebAuthnAuthentication)
		}

		// RBAC routes
		rbac := protected.Group("/rbac")
		{
			// Role management
			rbac.POST("/roles", rbacHandler.CreateRole)
			rbac.DELETE("/roles", rbacHandler.DeleteRole)
			rbac.POST("/roles/assign", rbacHandler.AssignRole)
			rbac.POST("/roles/remove", rbacHandler.RemoveRole)
			
			// Permission management
			rbac.POST("/permissions", rbacHandler.CreatePermission)
			rbac.POST("/permissions/assign", rbacHandler.AssignPermission)
			rbac.POST("/permissions/remove", rbacHandler.RemovePermission)
			
			// Access control
			rbac.POST("/permissions/check", rbacHandler.CheckPermission)
			rbac.POST("/users/roles", rbacHandler.GetUserRoles)
			rbac.POST("/roles/permissions", rbacHandler.GetRolePermissions)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Authentication Service is running",
		})
	})
}