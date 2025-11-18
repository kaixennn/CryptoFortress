package handlers

import (
	"net/http"

	"github.com/cryptofortress/backend/keymgmt/internal/services"
	"github.com/gin-gonic/gin"
)

// ReplicationHandler handles key replication HTTP requests
type ReplicationHandler struct {
	replicationService services.ReplicationService
}

// NewReplicationHandler creates a new replication handler
func NewReplicationHandler(replicationService services.ReplicationService) *ReplicationHandler {
	return &ReplicationHandler{
		replicationService: replicationService,
	}
}

// ReplicateKeyRequest represents the key replication request payload
type ReplicateKeyRequest struct {
	KeyID  string   `json:"key_id" binding:"required"`
	Regions []string `json:"regions" binding:"required"`
}

// ReplicateKeyResponse represents the key replication response payload
type ReplicateKeyResponse struct {
	Message string `json:"message"`
}

// ReplicateKey handles key replication requests
func (h *ReplicationHandler) ReplicateKey(c *gin.Context) {
	var req ReplicateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Replicate key
	err := h.replicationService.ReplicateKey(req.KeyID, req.Regions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to replicate key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, ReplicateKeyResponse{
		Message: "Key replicated successfully",
	})
}

// EnableCrossRegionReplicationRequest represents the cross-region replication enabling request payload
type EnableCrossRegionReplicationRequest struct {
	KeyID   string   `json:"key_id" binding:"required"`
	Regions []string `json:"regions" binding:"required"`
}

// EnableCrossRegionReplicationResponse represents the cross-region replication enabling response payload
type EnableCrossRegionReplicationResponse struct {
	Message string `json:"message"`
}

// EnableCrossRegionReplication handles cross-region replication enabling requests
func (h *ReplicationHandler) EnableCrossRegionReplication(c *gin.Context) {
	var req EnableCrossRegionReplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Enable cross-region replication
	err := h.replicationService.EnableCrossRegionReplication(req.KeyID, req.Regions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable cross-region replication"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, EnableCrossRegionReplicationResponse{
		Message: "Cross-region replication enabled successfully",
	})
}

// DisableCrossRegionReplicationRequest represents the cross-region replication disabling request payload
type DisableCrossRegionReplicationRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// DisableCrossRegionReplicationResponse represents the cross-region replication disabling response payload
type DisableCrossRegionReplicationResponse struct {
	Message string `json:"message"`
}

// DisableCrossRegionReplication handles cross-region replication disabling requests
func (h *ReplicationHandler) DisableCrossRegionReplication(c *gin.Context) {
	var req DisableCrossRegionReplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Disable cross-region replication
	err := h.replicationService.DisableCrossRegionReplication(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to disable cross-region replication"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, DisableCrossRegionReplicationResponse{
		Message: "Cross-region replication disabled successfully",
	})
}

// BackupKeyRequest represents the key backup request payload
type BackupKeyRequest struct {
	KeyID       string `json:"key_id" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

// BackupKeyResponse represents the key backup response payload
type BackupKeyResponse struct {
	Message string `json:"message"`
}

// BackupKey handles key backup requests
func (h *ReplicationHandler) BackupKey(c *gin.Context) {
	var req BackupKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Backup key
	err := h.replicationService.BackupKey(req.KeyID, req.Destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to backup key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, BackupKeyResponse{
		Message: "Key backup created successfully",
	})
}

// RestoreKeyRequest represents the key restoration request payload
type RestoreKeyRequest struct {
	BackupID    string `json:"backup_id" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

// RestoreKeyResponse represents the key restoration response payload
type RestoreKeyResponse struct {
	Message string `json:"message"`
}

// RestoreKey handles key restoration requests
func (h *ReplicationHandler) RestoreKey(c *gin.Context) {
	var req RestoreKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Restore key
	err := h.replicationService.RestoreKey(req.BackupID, req.Destination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, RestoreKeyResponse{
		Message: "Key restored successfully",
	})
}