package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the configuration for the key management service
type Config struct {
	Port              string
	DatabaseURL       string
	VaultAddr         string
	VaultToken        string
	KeyRotationPeriod int // in days
	ShamirThreshold   int
	ShamirShares      int
	ReplicationEnabled bool
	ReplicationRegions []string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	port := getEnv("KEYMGMT_SERVICE_PORT", "8082")
	
	rotationPeriod, err := strconv.Atoi(getEnv("KEY_ROTATION_PERIOD", "90"))
	if err != nil {
		return nil, fmt.Errorf("invalid KEY_ROTATION_PERIOD: %v", err)
	}
	
	threshold, err := strconv.Atoi(getEnv("SHAMIR_THRESHOLD", "2"))
	if err != nil {
		return nil, fmt.Errorf("invalid SHAMIR_THRESHOLD: %v", err)
	}
	
	shares, err := strconv.Atoi(getEnv("SHAMIR_SHARES", "3"))
	if err != nil {
		return nil, fmt.Errorf("invalid SHAMIR_SHARES: %v", err)
	}
	
	replicationEnabled, err := strconv.ParseBool(getEnv("REPLICATION_ENABLED", "false"))
	if err != nil {
		return nil, fmt.Errorf("invalid REPLICATION_ENABLED: %v", err)
	}
	
	// Parse replication regions
	regionsStr := getEnv("REPLICATION_REGIONS", "")
	var regions []string
	if regionsStr != "" {
		// Split by comma
		for _, region := range strings.Split(regionsStr, ",") {
			regions = append(regions, strings.TrimSpace(region))
		}
	}
	
	return &Config{
		Port:               port,
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		VaultAddr:          os.Getenv("VAULT_ADDR"),
		VaultToken:         os.Getenv("VAULT_TOKEN"),
		KeyRotationPeriod:  rotationPeriod,
		ShamirThreshold:    threshold,
		ShamirShares:       shares,
		ReplicationEnabled: replicationEnabled,
		ReplicationRegions: regions,
	}, nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}