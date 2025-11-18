package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the configuration for the audit & compliance service
type Config struct {
	Port                 string
	DatabaseURL          string
	AuditLogRetention    int // in days
	ComplianceStandards  []string
	SIEMEnabled          bool
	SIEMEndpoints        []string
	ImmutableAuditTrails bool
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	port := getEnv("AUDIT_SERVICE_PORT", "8083")
	
	retention, err := strconv.Atoi(getEnv("AUDIT_LOG_RETENTION", "365"))
	if err != nil {
		return nil, fmt.Errorf("invalid AUDIT_LOG_RETENTION: %v", err)
	}
	
	siemEnabled, err := strconv.ParseBool(getEnv("SIEM_ENABLED", "false"))
	if err != nil {
		return nil, fmt.Errorf("invalid SIEM_ENABLED: %v", err)
	}
	
	immutableAuditTrails, err := strconv.ParseBool(getEnv("IMMUTABLE_AUDIT_TRAILS", "true"))
	if err != nil {
		return nil, fmt.Errorf("invalid IMMUTABLE_AUDIT_TRAILS: %v", err)
	}
	
	// Parse compliance standards
	standardsStr := getEnv("COMPLIANCE_STANDARDS", "GDPR,HIPAA,SOC2")
	var standards []string
	if standardsStr != "" {
		for _, standard := range strings.Split(standardsStr, ",") {
			standards = append(standards, strings.TrimSpace(standard))
		}
	}
	
	// Parse SIEM endpoints
	endpointsStr := getEnv("SIEM_ENDPOINTS", "")
	var endpoints []string
	if endpointsStr != "" {
		for _, endpoint := range strings.Split(endpointsStr, ",") {
			endpoints = append(endpoints, strings.TrimSpace(endpoint))
		}
	}
	
	return &Config{
		Port:                 port,
		DatabaseURL:          os.Getenv("DATABASE_URL"),
		AuditLogRetention:    retention,
		ComplianceStandards:  standards,
		SIEMEnabled:          siemEnabled,
		SIEMEndpoints:        endpoints,
		ImmutableAuditTrails: immutableAuditTrails,
	}, nil
}

// getEnv retrieves environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}