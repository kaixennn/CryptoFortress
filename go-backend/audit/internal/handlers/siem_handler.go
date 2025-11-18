package handlers

import (
	"net/http"
	"time"

	"github.com/cryptofortress/backend/audit/internal/services"
	"github.com/gin-gonic/gin"
)

// SIEMHandler handles SIEM HTTP requests
type SIEMHandler struct {
	siemService services.SIEMService
}

// NewSIEMHandler creates a new SIEM handler
func NewSIEMHandler(siemService services.SIEMService) *SIEMHandler {
	return &SIEMHandler{
		siemService: siemService,
	}
}

// SendEventRequest represents the SIEM event sending request payload
type SendEventRequest struct {
	Source  string            `json:"source" binding:"required"`
	Type    string            `json:"type" binding:"required"`
	Severity string           `json:"severity" binding:"required"`
	Message string            `json:"message" binding:"required"`
	Data    map[string]string `json:"data,omitempty"`
}

// SendEventResponse represents the SIEM event sending response payload
type SendEventResponse struct {
	EventID string `json:"event_id"`
	Message string `json:"message"`
}

// SendEvent handles SIEM event sending requests
func (h *SIEMHandler) SendEvent(c *gin.Context) {
	var req SendEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create SIEM event
	event := services.SIEMEvent{
		Timestamp: time.Now(),
		Source:    req.Source,
		Type:      req.Type,
		Severity:  req.Severity,
		Message:   req.Message,
		Data:      req.Data,
	}

	// Send event
	err := h.siemService.SendEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SIEM event"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, SendEventResponse{
		EventID: event.ID,
		Message: "SIEM event sent successfully",
	})
}

// MonitorThreatsResponse represents the threat monitoring response payload
type MonitorThreatsResponse struct {
	Alerts  []ThreatAlertResponse `json:"alerts"`
	Message string                `json:"message"`
}

// ThreatAlertResponse represents a threat alert in the response
type ThreatAlertResponse struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Source    string `json:"source"`
	Type      string `json:"type"`
	Severity  string `json:"severity"`
	Message   string `json:"message"`
	Details   string `json:"details"`
}

// MonitorThreats handles threat monitoring requests
func (h *SIEMHandler) MonitorThreats(c *gin.Context) {
	// Monitor threats
	alerts, err := h.siemService.MonitorThreats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to monitor threats"})
		return
	}

	// Convert alerts to response format
	alertResponses := make([]ThreatAlertResponse, len(alerts))
	for i, alert := range alerts {
		alertResponses[i] = ThreatAlertResponse{
			ID:        alert.ID,
			Timestamp: alert.Timestamp.Format(time.RFC3339),
			Source:    alert.Source,
			Type:      alert.Type,
			Severity:  alert.Severity,
			Message:   alert.Message,
			Details:   alert.Details,
		}
	}

	// Return response
	c.JSON(http.StatusOK, MonitorThreatsResponse{
		Alerts:  alertResponses,
		Message: "Threat monitoring completed",
	})
}

// ConfigureEndpointRequest represents the SIEM endpoint configuration request payload
type ConfigureEndpointRequest struct {
	Endpoint string `json:"endpoint" binding:"required"`
}

// ConfigureEndpointResponse represents the SIEM endpoint configuration response payload
type ConfigureEndpointResponse struct {
	Message string `json:"message"`
}

// ConfigureEndpoint handles SIEM endpoint configuration requests
func (h *SIEMHandler) ConfigureEndpoint(c *gin.Context) {
	var req ConfigureEndpointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Configure endpoint
	err := h.siemService.ConfigureEndpoint(req.Endpoint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to configure SIEM endpoint"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, ConfigureEndpointResponse{
		Message: "SIEM endpoint configured successfully",
	})
}

// CreateAlertRuleRequest represents the alert rule creation request payload
type CreateAlertRuleRequest struct {
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description"`
	Condition   string            `json:"condition" binding:"required"`
	Severity    string            `json:"severity" binding:"required"`
	Enabled     bool              `json:"enabled"`
	Actions     []string          `json:"actions"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// CreateAlertRuleResponse represents the alert rule creation response payload
type CreateAlertRuleResponse struct {
	RuleID  string `json:"rule_id"`
	Message string `json:"message"`
}

// CreateAlertRule handles alert rule creation requests
func (h *SIEMHandler) CreateAlertRule(c *gin.Context) {
	var req CreateAlertRuleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create alert rule
	rule := services.AlertRule{
		Name:        req.Name,
		Description: req.Description,
		Condition:   req.Condition,
		Severity:    req.Severity,
		Enabled:     req.Enabled,
		Actions:     req.Actions,
		Metadata:    req.Metadata,
	}

	// Create rule
	err := h.siemService.CreateAlertRule(rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create alert rule"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CreateAlertRuleResponse{
		RuleID:  rule.ID,
		Message: "Alert rule created successfully",
	})
}