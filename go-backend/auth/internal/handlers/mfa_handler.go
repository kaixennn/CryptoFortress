package handlers

import (
	"net/http"

	"github.com/cryptofortress/backend/auth/internal/services"
	"github.com/gin-gonic/gin"
)

// MFAHandler handles multi-factor authentication HTTP requests
type MFAHandler struct {
	mfaService services.MFAService
}

// NewMFAHandler creates a new MFA handler
func NewMFAHandler(mfaService services.MFAService) *MFAHandler {
	return &MFAHandler{
		mfaService: mfaService,
	}
}

// EnableTOTPRequest represents the enable TOTP request payload
type EnableTOTPRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// EnableTOTPResponse represents the enable TOTP response payload
type EnableTOTPResponse struct {
	Secret string `json:"secret"`
	QRCode string `json:"qr_code"`
}

// EnableTOTP handles enabling TOTP for a user
func (h *MFAHandler) EnableTOTP(c *gin.Context) {
	var req EnableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate TOTP secret
	secret, err := h.mfaService.EnableTOTP(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate TOTP secret"})
		return
	}

	// In a real implementation, you would generate a QR code
	// For now, we'll just return the secret
	qrCode := "otpauth://totp/CryptoFortress:" + req.UserID + "?secret=" + secret + "&issuer=CryptoFortress"

	// Return response
	c.JSON(http.StatusOK, EnableTOTPResponse{
		Secret: secret,
		QRCode: qrCode,
	})
}

// VerifyTOTPRequest represents the verify TOTP request payload
type VerifyTOTPRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

// VerifyTOTP handles verifying a TOTP token
func (h *MFAHandler) VerifyTOTP(c *gin.Context) {
	var req VerifyTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify TOTP token
	valid := h.mfaService.VerifyTOTP(req.UserID, req.Token)
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid TOTP token"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "TOTP verification successful"})
}

// DisableTOTPRequest represents the disable TOTP request payload
type DisableTOTPRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// DisableTOTP handles disabling TOTP for a user
func (h *MFAHandler) DisableTOTP(c *gin.Context) {
	var req DisableTOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Disable TOTP
	err := h.mfaService.DisableTOTP(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to disable TOTP"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "TOTP disabled successfully"})
}

// RegisterWebAuthnRequest represents the register WebAuthn request payload
type RegisterWebAuthnRequest struct {
	UserID          string `json:"user_id" binding:"required"`
	CredentialName  string `json:"credential_name" binding:"required"`
}

// RegisterWebAuthn handles initiating WebAuthn credential registration
func (h *MFAHandler) RegisterWebAuthn(c *gin.Context) {
	var req RegisterWebAuthnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate WebAuthn registration options
	options, err := h.mfaService.RegisterWebAuthnCredential(req.UserID, req.CredentialName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate WebAuthn registration options"})
		return
	}

	// Return registration options
	c.Data(http.StatusOK, "application/json", options)
}

// VerifyWebAuthnRegistrationRequest represents the verify WebAuthn registration request payload
type VerifyWebAuthnRegistrationRequest struct {
	UserID               string `json:"user_id" binding:"required"`
	RegistrationResponse []byte `json:"registration_response" binding:"required"`
}

// VerifyWebAuthnRegistration handles completing WebAuthn credential registration
func (h *MFAHandler) VerifyWebAuthnRegistration(c *gin.Context) {
	var req VerifyWebAuthnRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify WebAuthn registration
	err := h.mfaService.VerifyWebAuthnRegistration(req.UserID, req.RegistrationResponse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify WebAuthn registration"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "WebAuthn registration successful"})
}

// AuthenticateWebAuthnRequest represents the authenticate WebAuthn request payload
type AuthenticateWebAuthnRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

// AuthenticateWebAuthn handles initiating WebAuthn authentication
func (h *MFAHandler) AuthenticateWebAuthn(c *gin.Context) {
	var req AuthenticateWebAuthnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate WebAuthn authentication options
	options, err := h.mfaService.AuthenticateWithWebAuthn(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate WebAuthn authentication options"})
		return
	}

	// Return authentication options
	c.Data(http.StatusOK, "application/json", options)
}

// VerifyWebAuthnAuthenticationRequest represents the verify WebAuthn authentication request payload
type VerifyWebAuthnAuthenticationRequest struct {
	UserID          string `json:"user_id" binding:"required"`
	AuthResponse    []byte `json:"auth_response" binding:"required"`
}

// VerifyWebAuthnAuthentication handles completing WebAuthn authentication
func (h *MFAHandler) VerifyWebAuthnAuthentication(c *gin.Context) {
	var req VerifyWebAuthnAuthenticationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify WebAuthn authentication
	err := h.mfaService.VerifyWebAuthnAuthentication(req.UserID, req.AuthResponse)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "WebAuthn authentication failed"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "WebAuthn authentication successful"})
}