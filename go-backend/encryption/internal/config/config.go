package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the configuration for the encryption service
type Config struct {
	Port              string
	DatabaseURL       string
	HSMEnabled        bool
	HSMAddress        string
	HSMToken          string
	VaultAddr         string
	VaultToken        string
	AESKeySize        int // in bits
	RSAKeySize        int // in bits
	DefaultAlgorithm  string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	port := getEnv("ENCRYPTION_SERVICE_PORT", "8081")
	
	aesKeySize, err := strconv.Atoi(getEnv("AES_KEY_SIZE", "256"))
	if err != nil {
		return nil, fmt.Errorf("invalid AES_KEY_SIZE: %v", err)
	}
	
	rsaKeySize, err := strconv.Atoi(getEnv("RSA_KEY_SIZE", "2048"))
	if err != nil {
		return nil, fmt.Errorf("invalid RSA_KEY_SIZE: %v", err)
	}
	
	hsmEnabled, err := strconv.ParseBool(getEnv("HSM_ENABLED", "false"))
	if err != nil {
		return nil, fmt.Errorf("invalid HSM_ENABLED: %v", err)
	}
	
	return &Config{
		Port:             port,
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		HSMEnabled:       hsmEnabled,
		HSMAddress:       os.Getenv("HSM_ADDRESS"),
		HSMToken:         os.Getenv("HSM_TOKEN"),
		VaultAddr:        os.Getenv("VAULT_ADDR"),
		VaultToken:       os.Getenv("VAULT_TOKEN"),
		AESKeySize:       aesKeySize,
		RSAKeySize:       rsaKeySize,
		DefaultAlgorithm: getEnv("DEFAULT_ALGORITHM", "AES-256-GCM"),
	}, nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}