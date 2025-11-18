package handlers

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/cryptofortress/backend/keymgmt/internal/services"
	"github.com/gin-gonic/gin"
)

// KeyHandler handles key management HTTP requests
type KeyHandler struct {
	keyService services.KeyService
}

// NewKeyHandler creates a new key handler
func NewKeyHandler(keyService services.KeyService) *KeyHandler {
	return &KeyHandler{
		keyService: keyService,
	}
}

// GenerateKeyRequest represents the key generation request payload
type GenerateKeyRequest struct {
	Algorithm string `json:"algorithm" binding:"required"`
	Size      int    `json:"size" binding:"required"`
}

// GenerateKeyResponse represents the key generation response payload
type GenerateKeyResponse struct {
	KeyID string `json:"key_id"`
	Key   string `json:"key"`
}

// GenerateKey handles key generation requests
func (h *KeyHandler) GenerateKey(c *gin.Context) {
	var req GenerateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate key
	keyID, key, err := h.keyService.GenerateKey(req.Algorithm, req.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate key"})
		return
	}

	// Encode key to base64
	keyB64 := base64.StdEncoding.EncodeToString(key)

	// Return response
	c.JSON(http.StatusOK, GenerateKeyResponse{
		KeyID: keyID,
		Key:   keyB64,
	})
}

// GenerateKeyPairRequest represents the key pair generation request payload
type GenerateKeyPairRequest struct {
	Algorithm string `json:"algorithm" binding:"required"`
	Size      int    `json:"size" binding:"required"`
}

// GenerateKeyPairResponse represents the key pair generation response payload
type GenerateKeyPairResponse struct {
	KeyID      string `json:"key_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// GenerateKeyPair handles key pair generation requests
func (h *KeyHandler) GenerateKeyPair(c *gin.Context) {
	var req GenerateKeyPairRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate key pair
	keyID, privateKey, publicKey, err := h.keyService.GenerateKeyPair(req.Algorithm, req.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate key pair"})
		return
	}

	// Encode keys to base64
	privateKeyB64 := base64.StdEncoding.EncodeToString(privateKey)
	publicKeyB64 := base64.StdEncoding.EncodeToString(publicKey)

	// Return response
	c.JSON(http.StatusOK, GenerateKeyPairResponse{
		KeyID:      keyID,
		PrivateKey: privateKeyB64,
		PublicKey:  publicKeyB64,
	})
}

// StoreKeyRequest represents the key storage request payload
type StoreKeyRequest struct {
	KeyID     string            `json:"key_id" binding:"required"`
	Key       string            `json:"key" binding:"required"`
	Algorithm string            `json:"algorithm" binding:"required"`
	Size      int               `json:"size"`
	Labels    map[string]string `json:"labels,omitempty"`
}

// StoreKeyResponse represents the key storage response payload
type StoreKeyResponse struct {
	Message string `json:"message"`
}

// StoreKey handles key storage requests
func (h *KeyHandler) StoreKey(c *gin.Context) {
	var req StoreKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Decode key from base64
	key, err := base64.StdEncoding.DecodeString(req.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key format"})
		return
	}

	// Create metadata
	metadata := services.KeyMetadata{
		Algorithm: req.Algorithm,
		Size:      req.Size,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Labels:    req.Labels,
	}

	// Store key
	err = h.keyService.StoreKey(req.KeyID, key, metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, StoreKeyResponse{
		Message: "Key stored successfully",
	})
}

// RetrieveKeyRequest represents the key retrieval request payload
type RetrieveKeyRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// RetrieveKeyResponse represents the key retrieval response payload
type RetrieveKeyResponse struct {
	Key      string              `json:"key"`
	Metadata RetrieveKeyMetadata `json:"metadata"`
}

// RetrieveKeyMetadata represents the metadata in key retrieval response
type RetrieveKeyMetadata struct {
	Algorithm   string            `json:"algorithm"`
	Size        int               `json:"size"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	ExpiresAt   *time.Time        `json:"expires_at,omitempty"`
	RevokedAt   *time.Time        `json:"revoked_at,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Description string            `json:"description,omitempty"`
	Version     int               `json:"version"`
}

// RetrieveKey handles key retrieval requests
func (h *KeyHandler) RetrieveKey(c *gin.Context) {
	var req RetrieveKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve key
	key, metadata, err := h.keyService.RetrieveKey(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve key"})
		return
	}

	// Encode key to base64
	keyB64 := base64.StdEncoding.EncodeToString(key)

	// Convert metadata
	var expiresAt *time.Time
	if metadata.ExpiresAt != nil {
		t := *metadata.ExpiresAt
		expiresAt = &t
	}

	var revokedAt *time.Time
	if metadata.RevokedAt != nil {
		t := *metadata.RevokedAt
		revokedAt = &t
	}

	// Return response
	c.JSON(http.StatusOK, RetrieveKeyResponse{
		Key: keyB64,
		Metadata: RetrieveKeyMetadata{
			Algorithm:   metadata.Algorithm,
			Size:        metadata.Size,
			CreatedAt:   metadata.CreatedAt,
			UpdatedAt:   metadata.UpdatedAt,
			ExpiresAt:   expiresAt,
			RevokedAt:   revokedAt,
			Labels:      metadata.Labels,
			Description: metadata.Description,
			Version:     metadata.Version,
		},
	})
}

// DeleteKeyRequest represents the key deletion request payload
type DeleteKeyRequest struct {
	KeyID string `json:"key_id" binding:"required"`
}

// DeleteKeyResponse represents the key deletion response payload
type DeleteKeyResponse struct {
	Message string `json:"message"`
}

// DeleteKey handles key deletion requests
func (h *KeyHandler) DeleteKey(c *gin.Context) {
	var req DeleteKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete key
	err := h.keyService.DeleteKey(req.KeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete key"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, DeleteKeyResponse{
		Message: "Key deleted successfully",
	})
}