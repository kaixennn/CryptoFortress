package services

import (
	"github.com/cryptofortress/backend/auth/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// Services holds references to all authentication services
type Services struct {
	Auth AuthService
	MFA  MFAService
	RBAC RBACService
}

// AuthService defines the interface for authentication operations
type AuthService interface {
	// JWT operations
	GenerateAccessToken(userID string, roles []string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(tokenString string) (*TokenClaims, error)
	ValidateRefreshToken(tokenString string) (*TokenClaims, error)
	RevokeRefreshToken(tokenString string) error
	
	// User authentication
	AuthenticateUser(username, password string) (*User, error)
	RegisterUser(username, email, password string) (*User, error)
	
	// OAuth2 operations
	InitiateOAuthFlow(provider string) (string, error)
	HandleOAuthCallback(provider, code string) (*User, error)
	
	// SAML operations
	ProcessSAMLResponse(samlResponse string) (*User, error)
	GenerateSAMLRequest(issuer, destination string) (string, error)
	
	// LDAP operations
	AuthenticateWithLDAP(username, password string) (*User, error)
}

// MFAService defines the interface for multi-factor authentication operations
type MFAService interface {
	// TOTP operations
	EnableTOTP(userID string) (string, error) // Returns secret key
	VerifyTOTP(userID, token string) bool
	DisableTOTP(userID string) error
	
	// WebAuthn operations
	RegisterWebAuthnCredential(userID, credentialName string) ([]byte, error) // Returns registration options
	VerifyWebAuthnRegistration(userID string, registrationResponse []byte) error
	AuthenticateWithWebAuthn(userID string) ([]byte, error) // Returns authentication options
	VerifyWebAuthnAuthentication(userID string, authResponse []byte) error
}

// RBACService defines the interface for role-based access control operations
type RBACService interface {
	// Role operations
	CreateRole(name, description string) error
	DeleteRole(name string) error
	AssignRoleToUser(userID, roleName string) error
	RemoveRoleFromUser(userID, roleName string) error
	
	// Permission operations
	CreatePermission(name, description string) error
	AssignPermissionToRole(roleName, permissionName string) error
	RemovePermissionFromRole(roleName, permissionName string) error
	
	// Access control
	CheckPermission(userID, permissionName string) (bool, error)
	GetUserRoles(userID string) ([]string, error)
	GetRolePermissions(roleName string) ([]string, error)
}

// User represents a user in the system
type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

// TokenClaims represents the claims in a JWT token
type TokenClaims struct {
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}