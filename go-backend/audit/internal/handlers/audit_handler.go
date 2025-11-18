package handlers

import (
	"net/http"
	"time"

	"github.com/cryptofortress/backend/audit/internal/services"
	"github.com/gin-gonic/gin"
)

// AuditHandler handles audit HTTP requests
type AuditHandler struct {
	auditService services.AuditService
}

// NewAuditHandler creates a new audit handler
func NewAuditHandler(auditService services.AuditService) *AuditHandler {
	return &AuditHandler{
		auditService: auditService,
	}
}

// LogEventRequest represents the audit event logging request payload
type LogEventRequest struct {
	UserID      string            `json:"user_id" binding:"required"`
	Action      string            `json:"action" binding:"required"`
	Resource    string            `json:"resource" binding:"required"`
	IPAddress   string            `json:"ip_address"`
	UserAgent   string            `json:"user_agent"`
	Success     bool              `json:"success"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// LogEventResponse represents the audit event logging response payload
type LogEventResponse struct {
	Message string `json:"message"`
}

// LogEvent handles audit event logging requests
func (h *AuditHandler) LogEvent(c *gin.Context) {
	var req LogEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create audit event
	event := services.AuditEvent{
		Timestamp:   time.Now(),
		UserID:      req.UserID,
		Action:      req.Action,
		Resource:    req.Resource,
		IPAddress:   req.IPAddress,
		UserAgent:   req.UserAgent,
		Success:     req.Success,
		Description: req.Description,
		Metadata:    req.Metadata,
	}

	// Log event
	err := h.auditService.LogEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log audit event"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, LogEventResponse{
		Message: "Audit event logged successfully",
	})
}

// GetAuditTrailRequest represents the audit trail retrieval request payload
type GetAuditTrailRequest struct {
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	UserID    string     `json:"user_id,omitempty"`
	Action    string     `json:"action,omitempty"`
	Resource  string     `json:"resource,omitempty"`
	Success   *bool      `json:"success,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
}

// GetAuditTrailResponse represents the audit trail retrieval response payload
type GetAuditTrailResponse struct {
	Events []AuditEventResponse `json:"events"`
}

// AuditEventResponse represents an audit event in the response
type AuditEventResponse struct {
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
	Signature   string            `json:"signature,omitempty"`
}

// GetAuditTrail handles audit trail retrieval requests
func (h *AuditHandler) GetAuditTrail(c *gin.Context) {
	var req GetAuditTrailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create filter
	filter := services.AuditFilter{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserID:    req.UserID,
		Action:    req.Action,
		Resource:  req.Resource,
		Success:   req.Success,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	// Get audit trail
	events, err := h.auditService.GetAuditTrail(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve audit trail"})
		return
	}

	// Convert events to response format
	eventResponses := make([]AuditEventResponse, len(events))
	for i, event := range events {
		eventResponses[i] = AuditEventResponse{
			ID:          event.ID,
			Timestamp:   event.Timestamp,
			UserID:      event.UserID,
			Action:      event.Action,
			Resource:    event.Resource,
			IPAddress:   event.IPAddress,
			UserAgent:   event.UserAgent,
			Success:     event.Success,
			Description: event.Description,
			Metadata:    event.Metadata,
			Signature:   event.Signature,
		}
	}

	// Return response
	c.JSON(http.StatusOK, GetAuditTrailResponse{
		Events: eventResponses,
	})
}

// CreateImmutableTrailRequest represents the immutable trail creation request payload
type CreateImmutableTrailRequest struct {
	UserID      string            `json:"user_id" binding:"required"`
	Action      string            `json:"action" binding:"required"`
	Resource    string            `json:"resource" binding:"required"`
	IPAddress   string            `json:"ip_address"`
	UserAgent   string            `json:"user_agent"`
	Success     bool              `json:"success"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// CreateImmutableTrailResponse represents the immutable trail creation response payload
type CreateImmutableTrailResponse struct {
	TrailID   string `json:"trail_id"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

// CreateImmutableTrail handles immutable trail creation requests
func (h *AuditHandler) CreateImmutableTrail(c *gin.Context) {
	var req CreateImmutableTrailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create audit event
	event := services.AuditEvent{
		UserID:      req.UserID,
		Action:      req.Action,
		Resource:    req.Resource,
		IPAddress:   req.IPAddress,
		UserAgent:   req.UserAgent,
		Success:     req.Success,
		Description: req.Description,
		Metadata:    req.Metadata,
	}

	// Create immutable trail
	trailID, err := h.auditService.CreateImmutableTrail(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create immutable trail"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CreateImmutableTrailResponse{
		TrailID:   trailID,
		Message:   "Immutable trail created successfully",
		Signature: event.Signature,
	})
}

// VerifyTrailIntegrityRequest represents the trail integrity verification request payload
type VerifyTrailIntegrityRequest struct {
	TrailID string `json:"trail_id" binding:"required"`
}

// VerifyTrailIntegrityResponse represents the trail integrity verification response payload
type VerifyTrailIntegrityResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}

// VerifyTrailIntegrity handles trail integrity verification requests
func (h *AuditHandler) VerifyTrailIntegrity(c *gin.Context) {
	var req VerifyTrailIntegrityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify trail integrity
	valid, err := h.auditService.VerifyTrailIntegrity(req.TrailID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify trail integrity"})
		return
	}

	// Return response
	message := "Trail integrity verified"
	if !valid {
		message = "Trail integrity verification failed"
	}

	c.JSON(http.StatusOK, VerifyTrailIntegrityResponse{
		Valid:   valid,
		Message: message,
	})
}