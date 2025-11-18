package handlers

import (
	"net/http"
	"time"

	"github.com/cryptofortress/backend/keymgmt/internal/services"
	"github.com/gin-gonic/gin"
)

// RotationHandler handles key rotation HTTP requests
type RotationHandler struct {
	rotationService services.RotationService
}

// NewRotationHandler creates a new rotation handler
func NewRotationHandler(rotationService services.RotationService) *RotationHandler {
	return &RotationHandler{
		rotationService: rotationService,
	}
}

// RotateKeyRequest represents the key rotation request payload
type RotateKeyRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// RotateKeyResponse represents the key rotation response payload
type RotateKeyResponse struct {
	NewKeyID string `json:"new_key_id"`
	Message  string `json:"message"`
}

// RotateKey handles key rotation requests
func (h *RotationHandler) RotateKey(c *gin.Context) {
	var req RotateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Rotate key
	newKeyID, err := h.rotationService.RotateKey(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to rotate key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, RotateKeyResponse{
		NewKeyID: newKeyID,
		Message:  "Key rotated successfully",
	})
}

// ScheduleRotationRequest represents the rotation scheduling request payload
type ScheduleRotationRequest struct {
	KeyID    string `json:"key_id" binding:"required"`
	Interval string `json:"interval" binding:"required"` // e.g., "24h", "720h"
}

// ScheduleRotationResponse represents the rotation scheduling response payload
type ScheduleRotationResponse struct {
	Message string `json:"message"`
}

// ScheduleRotation handles rotation scheduling requests
func (h *RotationHandler) ScheduleRotation(c *gin.Context) {
	var req ScheduleRotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse interval
	interval, err := time.ParseDuration(req.Interval)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval format"})
		return
	}

	// Schedule rotation
	err = h.rotationService.ScheduleRotation(req.KeyID, interval)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule rotation"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, ScheduleRotationResponse{
		Message: "Rotation scheduled successfully",
	})
}

// CancelRotationRequest represents the rotation cancellation request payload
type CancelRotationRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// CancelRotationResponse represents the rotation cancellation response payload
type CancelRotationResponse struct {
	Message string `json:"message"`
}

// CancelRotation handles rotation cancellation requests
func (h *RotationHandler) CancelRotation(c *gin.Context) {
	var req CancelRotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cancel rotation
	err := h.rotationService.CancelRotation(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel rotation"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CancelRotationResponse{
		Message: "Rotation cancelled successfully",
	})
}

// GetRotationScheduleRequest represents the rotation schedule retrieval request payload
type GetRotationScheduleRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// GetRotationScheduleResponse represents the rotation schedule retrieval response payload
type GetRotationScheduleResponse struct {
	KeyID      string    `json:"key_id"`
	Interval   string    `json:"interval"`
	NextRotate time.Time `json:"next_rotate"`
	Enabled    bool      `json:"enabled"`
}

// GetRotationSchedule handles rotation schedule retrieval requests
func (h *RotationHandler) GetRotationSchedule(c *gin.Context) {
	var req GetRotationScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get rotation schedule
	schedule, err := h.rotationService.GetRotationSchedule(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rotation schedule"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, GetRotationScheduleResponse{
		KeyID:      schedule.KeyID,
		Interval:   schedule.Interval.String(),
		NextRotate: schedule.NextRotate,
		Enabled:    schedule.Enabled,
	})
}

// EnableAutoRotationRequest represents the auto rotation enabling request payload
type EnableAutoRotationRequest struct {
	KeyID  string `json:"key_id" binding:"required"`
	Period string `json:"period" binding:"required"` // e.g., "24h", "720h"
}

// EnableAutoRotationResponse represents the auto rotation enabling response payload
type EnableAutoRotationResponse struct {
	Message string `json:"message"`
}

// EnableAutoRotation handles auto rotation enabling requests
func (h *RotationHandler) EnableAutoRotation(c *gin.Context) {
	var req EnableAutoRotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse period
	period, err := time.ParseDuration(req.Period)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid period format"})
		return
	}

	// Enable auto rotation
	err = h.rotationService.EnableAutoRotation(req.KeyID, period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable auto rotation"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, EnableAutoRotationResponse{
		Message: "Auto rotation enabled successfully",
	})
}

// DisableAutoRotationRequest represents the auto rotation disabling request payload
type DisableAutoRotationRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// DisableAutoRotationResponse represents the auto rotation disabling response payload
type DisableAutoRotationResponse struct {
	Message string `json:"message"`
}

// DisableAutoRotation handles auto rotation disabling requests
func (h *RotationHandler) DisableAutoRotation(c *gin.Context) {
	var req DisableAutoRotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Disable auto rotation
	err := h.rotationService.DisableAutoRotation(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to disable auto rotation"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, DisableAutoRotationResponse{
		Message: "Auto rotation disabled successfully",
	})
}