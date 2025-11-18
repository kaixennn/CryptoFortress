package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/cryptofortress/backend/encryption/internal/services"
	"github.com/gin-gonic/gin"
)

// EncryptionHandler handles encryption-related HTTP requests
type EncryptionHandler struct {
	encryptionService services.EncryptionService
}

// NewEncryptionHandler creates a new encryption handler
func NewEncryptionHandler(encryptionService services.EncryptionService) *EncryptionHandler {
	return &EncryptionHandler{
		encryptionService: encryptionService,
	}
}

// EncryptRequest represents the encryption request payload
type EncryptRequest struct {
	Plaintext string `json:"plaintext" binding:"required"`
	Key       string `json:"key" binding:"required"`
	Algorithm string `json:"algorithm" binding:"required"`
}

// EncryptResponse represents the encryption response payload
type EncryptResponse struct {
	Ciphertext string `json:"ciphertext"`
	Nonce      string `json:"nonce,omitempty"`
}

// Encrypt handles encryption requests
func (h *EncryptionHandler) Encrypt(c *gin.Context) {
	var req EncryptRequest
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

	// Convert plaintext to bytes
	plaintext := []byte(req.Plaintext)

	// Encrypt based on algorithm
	var ciphertext []byte
	var nonce []byte

	switch req.Algorithm {
	case "AES-256-GCM":
		// Generate nonce
		nonce, err = h.encryptionService.GenerateNonce(12) // 12 bytes for GCM
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
			return
		}

		// Encrypt
		ciphertext, err = h.encryptionService.EncryptAES256GCM(plaintext, key, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
			return
		}
	case "ChaCha20-Poly1305":
		// Generate nonce
		nonce, err = h.encryptionService.GenerateNonce(12) // 12 bytes for ChaCha20-Poly1305
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
			return
		}

		// Encrypt
		ciphertext, err = h.encryptionService.EncryptChaCha20Poly1305(plaintext, key, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Encryption failed"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported algorithm"})
		return
	}

	// Encode ciphertext and nonce to base64
	ciphertextB64 := base64.StdEncoding.EncodeToString(ciphertext)
	nonceB64 := base64.StdEncoding.EncodeToString(nonce)

	// Return response
	c.JSON(http.StatusOK, EncryptResponse{
		Ciphertext: ciphertextB64,
		Nonce:      nonceB64,
	})
}

// DecryptRequest represents the decryption request payload
type DecryptRequest struct {
	Ciphertext string `json:"ciphertext" binding:"required"`
	Key        string `json:"key" binding:"required"`
	Nonce      string `json:"nonce" binding:"required"`
	Algorithm  string `json:"algorithm" binding:"required"`
}

// DecryptResponse represents the decryption response payload
type DecryptResponse struct {
	Plaintext string `json:"plaintext"`
}

// Decrypt handles decryption requests
func (h *EncryptionHandler) Decrypt(c *gin.Context) {
	var req DecryptRequest
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

	// Decode ciphertext from base64
	ciphertext, err := base64.StdEncoding.DecodeString(req.Ciphertext)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ciphertext format"})
		return
	}

	// Decode nonce from base64
	nonce, err := base64.StdEncoding.DecodeString(req.Nonce)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid nonce format"})
		return
	}

	// Decrypt based on algorithm
	var plaintext []byte

	switch req.Algorithm {
	case "AES-256-GCM":
		plaintext, err = h.encryptionService.DecryptAES256GCM(ciphertext, key, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
			return
		}
	case "ChaCha20-Poly1305":
		plaintext, err = h.encryptionService.DecryptChaCha20Poly1305(ciphertext, key, nonce)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Decryption failed"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported algorithm"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, DecryptResponse{
		Plaintext: string(plaintext),
	})
}

// GenerateKeyRequest represents the key generation request payload
type GenerateKeyRequest struct {
	Algorithm string `json:"algorithm" binding:"required"`
}

// GenerateKeyResponse represents the key generation response payload
type GenerateKeyResponse struct {
	Key string `json:"key"`
}

// GenerateKey handles key generation requests
func (h *EncryptionHandler) GenerateKey(c *gin.Context) {
	var req GenerateKeyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate key based on algorithm
	var key []byte
	var err error

	switch req.Algorithm {
	case "AES-256-GCM", "ChaCha20-Poly1305":
		key, err = h.encryptionService.GenerateAESKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate key"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported algorithm"})
		return
	}

	// Encode key to base64
	keyB64 := base64.StdEncoding.EncodeToString(key)

	// Return response
	c.JSON(http.StatusOK, GenerateKeyResponse{
		Key: keyB64,
	})
}