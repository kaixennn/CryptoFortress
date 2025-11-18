package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the configuration for the authentication service
type Config struct {
	Port              string
	JWTSecret         string
	RefreshTokenTTL   int // in hours
	AccessTokenTTL    int // in minutes
	DatabaseURL       string
	LDAPServer        string
	LDAPPort          int
	OAuthClientID     string
	OAuthClientSecret string
	SAMLEntityID      string
	VaultAddr         string
	VaultToken        string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	port := getEnv("AUTH_SERVICE_PORT", "8080")
	
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET environment variable is required")
	}
	
	refreshTokenTTL, err := strconv.Atoi(getEnv("REFRESH_TOKEN_TTL", "720")) // 30 days default
	if err != nil {
		return nil, fmt.Errorf("invalid REFRESH_TOKEN_TTL: %v", err)
	}
	
	accessTokenTTL, err := strconv.Atoi(getEnv("ACCESS_TOKEN_TTL", "15")) // 15 minutes default
	if err != nil {
		return nil, fmt.Errorf("invalid ACCESS_TOKEN_TTL: %v", err)
	}
	
	ldapPort, err := strconv.Atoi(getEnv("LDAP_PORT", "389"))
	if err != nil {
		return nil, fmt.Errorf("invalid LDAP_PORT: %v", err)
	}
	
	return &Config{
		Port:              port,
		JWTSecret:         jwtSecret,
		RefreshTokenTTL:   refreshTokenTTL,
		AccessTokenTTL:    accessTokenTTL,
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		LDAPServer:        os.Getenv("LDAP_SERVER"),
		LDAPPort:          ldapPort,
		OAuthClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		OAuthClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		SAMLEntityID:      os.Getenv("SAML_ENTITY_ID"),
		VaultAddr:         os.Getenv("VAULT_ADDR"),
		VaultToken:        os.Getenv("VAULT_TOKEN"),
	}, nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}