package services

import (
	"time"

	"github.com/cryptofortress/backend/audit/internal/config"
)

// Services holds references to all audit & compliance services
type Services struct {
	Audit      AuditService
	Compliance ComplianceService
	SIEM       SIEMService
}

// AuditService defines the interface for audit operations
type AuditService interface {
	// Audit trail management
	LogEvent(event AuditEvent) error
	GetAuditTrail(filter AuditFilter) ([]AuditEvent, error)
	ExportAuditTrail(format string, filter AuditFilter) ([]byte, error)
	
	// Immutable audit trails
	CreateImmutableTrail(event AuditEvent) (string, error)
	VerifyTrailIntegrity(trailID string) (bool, error)
	
	// Audit log retention
	SetRetentionPeriod(days int) error
	PurgeExpiredLogs() error
}

// ComplianceService defines the interface for compliance operations
type ComplianceService interface {
	// Compliance reporting
	GenerateReport(standard string, period time.Time) (*ComplianceReport, error)
	GetComplianceStatus(standard string) (*ComplianceStatus, error)
	ListComplianceStandards() ([]string, error)
	
	// GDPR specific operations
	HandleDataSubjectRequest(request DataSubjectRequest) error
	GenerateDataInventory() (*DataInventory, error)
	
	// HIPAA specific operations
	GenerateHIPAAReport() (*HIPAAReport, error)
	AuditPHIAccess(userID string) ([]PHIEvent, error)
	
	// SOC2 specific operations
	GenerateSOC2Report() (*SOC2Report, error)
	AuditSystemSecurity() ([]SecurityEvent, error)
}

// SIEMService defines the interface for SIEM operations
type SIEMService interface {
	// Real-time monitoring
	SendEvent(event SIEMEvent) error
	MonitorThreats() ([]ThreatAlert, error)
	
	// SIEM integration
	ConfigureEndpoint(endpoint string) error
	RemoveEndpoint(endpoint string) error
	SendToAllEndpoints(event SIEMEvent) error
	
	// Alerting
	CreateAlertRule(rule AlertRule) error
	EvaluateAlertRules(event SIEMEvent) ([]Alert, error)
}

// AuditEvent represents an audit event
type AuditEvent struct {
	ID          string            `json:"id"`
	Timestamp   time.Time         `json:"timestamp"`
	UserID      string            `json:"user_id"`
	Action      string            `json:"action"`
	Resource    string            `json:"resource"`
	IPAddress   string            `json:"ip_address"`
	UserAgent   string            `json:"user_agent"`
	Success     bool              `json:"success"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Signature   string            `json:"signature,omitempty"` // For immutable trails
}

// AuditFilter represents filters for audit trail queries
type AuditFilter struct {
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	UserID     string     `json:"user_id,omitempty"`
	Action     string     `json:"action,omitempty"`
	Resource   string     `json:"resource,omitempty"`
	Success    *bool      `json:"success,omitempty"`
	Limit      int        `json:"limit,omitempty"`
	Offset     int        `json:"offset,omitempty"`
}

// ComplianceReport represents a compliance report
type ComplianceReport struct {
	ID        string    `json:"id"`
	Standard  string    `json:"standard"`
	Period    time.Time `json:"period"`
	Generated time.Time `json:"generated"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
}

// ComplianceStatus represents compliance status
type ComplianceStatus struct {
	Standard     string    `json:"standard"`
	Compliant    bool      `json:"compliant"`
	LastChecked  time.Time `json:"last_checked"`
	Issues       []string  `json:"issues,omitempty"`
	NextAudit    time.Time `json:"next_audit"`
}

// DataSubjectRequest represents a GDPR data subject request
type DataSubjectRequest struct {
	ID           string    `json:"id"`
	RequestType  string    `json:"request_type"` // "access", "rectification", "erasure", "restriction", "portability"
	RequesterID  string    `json:"requester_id"`
	RequestedAt  time.Time `json:"requested_at"`
	ProcessedAt  *time.Time `json:"processed_at,omitempty"`
	Status       string    `json:"status"` // "pending", "processing", "completed", "rejected"
	Data         string    `json:"data,omitempty"`
}

// DataInventory represents a data inventory for GDPR
type DataInventory struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Datasets  []DatasetInfo  `json:"datasets"`
}

// DatasetInfo represents information about a dataset
type DatasetInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    string            `json:"category"`
	Retention   string            `json:"retention"`
	Controllers []string          `json:"controllers"`
	Processors  []string          `json:"processors"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// HIPAAReport represents a HIPAA compliance report
type HIPAAReport struct {
	ID        string        `json:"id"`
	Period    time.Time     `json:"period"`
	Generated time.Time     `json:"generated"`
	Entities  []HIPAAEntity `json:"entities"`
}

// HIPAAEntity represents a HIPAA-covered entity
type HIPAAEntity struct {
	Name       string       `json:"name"`
	Type       string       `json:"type"` // "covered_entity", "business_associate"
	PHIEvents  []PHIEvent   `json:"phi_events"`
	Audits     []HIPAAAudit `json:"audits"`
}

// PHIEvent represents a protected health information event
type PHIEvent struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	UserID      string    `json:"user_id"`
	PatientID   string    `json:"patient_id"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	Authorized  bool      `json:"authorized"`
}

// HIPAAAudit represents a HIPAA audit
type HIPAAAudit struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	Type      string    `json:"type"`
	Findings  []string  `json:"findings"`
	Resolved  bool      `json:"resolved"`
}

// SOC2Report represents a SOC2 compliance report
type SOC2Report struct {
	ID        string         `json:"id"`
	Period    time.Time      `json:"period"`
	Generated time.Time      `json:"generated"`
	TrustServices []TrustService `json:"trust_services"`
}

// TrustService represents a SOC2 trust service
type TrustService struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Controls    []SecurityControl `json:"controls"`
}

// SecurityControl represents a security control
type SecurityControl struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      string   `json:"status"` // "implemented", "partially_implemented", "not_implemented"
	Evidence    []string `json:"evidence,omitempty"`
}

// SecurityEvent represents a security event for SOC2
type SecurityEvent struct {
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Type        string    `json:"type"`
	Severity    string    `json:"severity"`
	Description string    `json:"description"`
	Remediated  bool      `json:"remediated"`
}

// SIEMEvent represents a SIEM event
type SIEMEvent struct {
	ID        string            `json:"id"`
	Timestamp time.Time         `json:"timestamp"`
	Source    string            `json:"source"`
	Type      string            `json:"type"`
	Severity  string            `json:"severity"`
	Message   string            `json:"message"`
	Data      map[string]string `json:"data,omitempty"`
}

// ThreatAlert represents a threat alert
type ThreatAlert struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	Details   string    `json:"details"`
}

// AlertRule represents an alert rule
type AlertRule struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Condition   string            `json:"condition"`
	Severity    string            `json:"severity"`
	Enabled     bool              `json:"enabled"`
	Actions     []string          `json:"actions"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// Alert represents an alert
type Alert struct {
	ID        string    `json:"id"`
	RuleID    string    `json:"rule_id"`
	Timestamp time.Time `json:"timestamp"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	Event     SIEMEvent `json:"event"`
}