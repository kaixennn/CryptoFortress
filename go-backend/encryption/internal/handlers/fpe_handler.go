package handlers

import (
	"encoding/base64"
	"net/http"

	"github.com/cryptofortress/backend/encryption/internal/services"
	"github.com/gin-gonic/gin"
)

// FPEHandler handles format-preserving encryption HTTP requests
type FPEHandler struct {
	fpeService services.FPEService
}

// NewFPEHandler creates a new FPE handler
func NewFPEHandler(fpeService services.FPEService) *FPEHandler {
	return &FPEHandler{
		fpeService: fpeService,
	}
}

// FPEEncryptRequest represents the FPE encryption request payload
type FPEEncryptRequest struct {
	Plaintext string `json:"plaintext" binding:"required"`
	Key       string `json:"key" binding:"required"`
	Tweak     string `json:"tweak,omitempty"`
}

// FPEEncryptResponse represents the FPE encryption response payload
type FPEEncryptResponse struct {
	Ciphertext string `json:"ciphertext"`
}

// FPEEncrypt handles FPE encryption requests
func (h *FPEHandler) FPEEncrypt(c *gin.Context) {
	var req FPEEncryptRequest
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

	// Decode tweak from base64 (if provided)
	var tweak []byte
	if req.Tweak != "" {
		tweak, err = base64.StdEncoding.DecodeString(req.Tweak)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweak format"})
			return
		}
	}

	// Encrypt
	ciphertext, err := h.fpeService.EncryptFPE(req.Plaintext, key, tweak)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FPE encryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, FPEEncryptResponse{
		Ciphertext: ciphertext,
	})
}

// FPEDecryptRequest represents the FPE decryption request payload
type FPEDecryptRequest struct {
	Ciphertext string `json:"ciphertext" binding:"required"`
	Key        string `json:"key" binding:"required"`
	Tweak      string `json:"tweak,omitempty"`
}

// FPEDecryptResponse represents the FPE decryption response payload
type FPEDecryptResponse struct {
	Plaintext string `json:"plaintext"`
}

// FPEDecrypt handles FPE decryption requests
func (h *FPEHandler) FPEDecrypt(c *gin.Context) {
	var req FPEDecryptRequest
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

	// Decode tweak from base64 (if provided)
	var tweak []byte
	if req.Tweak != "" {
		tweak, err = base64.StdEncoding.DecodeString(req.Tweak)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tweak format"})
			return
		}
	}

	// Decrypt
	plaintext, err := h.fpeService.DecryptFPE(req.Ciphertext, key, tweak)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "FPE decryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, FPEDecryptResponse{
		Plaintext: plaintext,
	})
}

// CreditCardEncryptRequest represents the credit card encryption request payload
type CreditCardEncryptRequest struct {
	CardNumber string `json:"card_number" binding:"required"`
	Key        string `json:"key" binding:"required"`
}

// CreditCardEncryptResponse represents the credit card encryption response payload
type CreditCardEncryptResponse struct {
	EncryptedCard string `json:"encrypted_card"`
}

// EncryptCreditCard handles credit card encryption requests
func (h *FPEHandler) EncryptCreditCard(c *gin.Context) {
	var req CreditCardEncryptRequest
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

	// Encrypt credit card
	encryptedCard, err := h.fpeService.EncryptCreditCard(req.CardNumber, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Credit card encryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CreditCardEncryptResponse{
		EncryptedCard: encryptedCard,
	})
}

// CreditCardDecryptRequest represents the credit card decryption request payload
type CreditCardDecryptRequest struct {
	EncryptedCard string `json:"encrypted_card" binding:"required"`
	Key           string `json:"key" binding:"required"`
}

// CreditCardDecryptResponse represents the credit card decryption response payload
type CreditCardDecryptResponse struct {
	CardNumber string `json:"card_number"`
}

// DecryptCreditCard handles credit card decryption requests
func (h *FPEHandler) DecryptCreditCard(c *gin.Context) {
	var req CreditCardDecryptRequest
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

	// Decrypt credit card
	cardNumber, err := h.fpeService.DecryptCreditCard(req.EncryptedCard, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Credit card decryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, CreditCardDecryptResponse{
		CardNumber: cardNumber,
	})
}

// SSNEncryptRequest represents the SSN encryption request payload
type SSNEncryptRequest struct {
	SSN string `json:"ssn" binding:"required"`
	Key string `json:"key" binding:"required"`
}

// SSNEncryptResponse represents the SSN encryption response payload
type SSNEncryptResponse struct {
	EncryptedSSN string `json:"encrypted_ssn"`
}

// EncryptSSN handles SSN encryption requests
func (h *FPEHandler) EncryptSSN(c *gin.Context) {
	var req SSNEncryptRequest
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

	// Encrypt SSN
	encryptedSSN, err := h.fpeService.EncryptSSN(req.SSN, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SSN encryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, SSNEncryptResponse{
		EncryptedSSN: encryptedSSN,
	})
}

// SSNDecryptRequest represents the SSN decryption request payload
type SSNDecryptRequest struct {
	EncryptedSSN string `json:"encrypted_ssn" binding:"required"`
	Key          string `json:"key" binding:"required"`
}

// SSNDecryptResponse represents the SSN decryption response payload
type SSNDecryptResponse struct {
	SSN string `json:"ssn"`
}

// DecryptSSN handles SSN decryption requests
func (h *FPEHandler) DecryptSSN(c *gin.Context) {
	var req SSNDecryptRequest
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

	// Decrypt SSN
	ssn, err := h.fpeService.DecryptSSN(req.EncryptedSSN, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "SSN decryption failed"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, SSNDecryptResponse{
		SSN: ssn,
	})
}