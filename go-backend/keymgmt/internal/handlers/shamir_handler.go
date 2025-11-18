package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/cryptofortress/backend/keymgmt/internal/services"
	"github.com/gin-gonic/gin"
)

// ShamirHandler handles Shamir's Secret Sharing HTTP requests
type ShamirHandler struct {
	shamirService services.ShamirService
}

// NewShamirHandler creates a new Shamir handler
func NewShamirHandler(shamirService services.ShamirService) *ShamirHandler {
	return &ShamirHandler{
		shamirService: shamirService,
	}
}

// SplitSecretRequest represents the secret splitting request payload
type SplitSecretRequest struct {
	Secret    string `json:"secret" binding:"required"`
	Threshold int    `json:"threshold" binding:"required"`
	Shares    int    `json:"shares" binding:"required"`
}

// SplitSecretResponse represents the secret splitting response payload
type SplitSecretResponse struct {
	ShareStrings []string `json:"shares"`
}

// SplitSecret handles secret splitting requests
func (h *ShamirHandler) SplitSecret(c *gin.Context) {
	var req SplitSecretRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert secret to bytes
	secret := []byte(req.Secret)

	// Split secret
	shareBytes, err := h.shamirService.SplitSecret(secret, req.Threshold, req.Shares)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to split secret"})
		return
	}

	// Convert shares to base64 strings
	shareStrings := make([]string, len(shareBytes))
	for i, share := range shareBytes {
		shareStrings[i] = base64.StdEncoding.EncodeToString(share)
	}

	// Return response
	c.JSON(http.StatusOK, SplitSecretResponse{
		ShareStrings: shareStrings,
	})
}

// CombineSharesRequest represents the share combination request payload
type CombineSharesRequest struct {
	Shares []string `json:"shares" binding:"required"`
}

// CombineSharesResponse represents the share combination response payload
type CombineSharesResponse struct {
	Secret string `json:"secret"`
}

// CombineShares handles share combination requests
func (h *ShamirHandler) CombineShares(c *gin.Context) {
	var req CombineSharesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert share strings to bytes
	shareBytes := make([][]byte, len(req.Shares))
	for i, shareStr := range req.Shares {
		share, err := base64.StdEncoding.DecodeString(shareStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid share format"})
			return
		}
		shareBytes[i] = share
	}

	// Combine shares
	secret, err := h.shamirService.CombineShares(shareBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to combine shares"})
		return
	}

	// Convert secret to string
	secretStr := string(secret)

	// Return response
	c.JSON(http.StatusOK, CombineSharesResponse{
		Secret: secretStr,
	})
}

// DistributeKeyRequest represents the key distribution request payload
type DistributeKeyRequest struct {
	KeyID     string   `json:"key_id" binding:"required"`
	Threshold int      `json:"threshold" binding:"required"`
	Shares    int      `json:"shares" binding:"required"`
	Recipients []string `json:"recipients" binding:"required"`
}

// DistributeKeyResponse represents the key distribution response payload
type DistributeKeyResponse struct {
	Message string `json:"message"`
}

// DistributeKey handles key distribution requests
func (h *ShamirHandler) DistributeKey(c *gin.Context) {
	var req DistributeKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Distribute key
	err := h.shamirService.DistributeKey(req.KeyID, req.Threshold, req.Shares, req.Recipients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to distribute key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, DistributeKeyResponse{
		Message: "Key distributed successfully",
	})
}

// RecoverKeyRequest represents the key recovery request payload
type RecoverKeyRequest struct {
	Shares []string `json:"shares" binding:"required"`
}

// RecoverKeyResponse represents the key recovery response payload
type RecoverKeyResponse struct {
	KeyID string `json:"key_id"`
	Key   string `json:"key"`
}

// RecoverKey handles key recovery requests
func (h *ShamirHandler) RecoverKey(c *gin.Context) {
	var req RecoverKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert share strings to bytes
	shareBytes := make([][]byte, len(req.Shares))
	for i, shareStr := range req.Shares {
		share, err := base64.StdEncoding.DecodeString(shareStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid share format"})
			return
		}
		shareBytes[i] = share
	}

	// Recover key
	keyID, key, err := h.shamirService.RecoverKey(shareBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to recover key"})
		return
	}

	// Encode key to base64
	keyB64 := base64.StdEncoding.EncodeToString(key)

	// Return response
	c.JSON(http.StatusOK, RecoverKeyResponse{
		KeyID: keyID,
		Key:   keyB64,
	})
}