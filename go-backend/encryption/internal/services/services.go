package services

import (
	"github.com/cryptofortress/backend/encryption/internal/config"
)

// Services holds references to all encryption services
type Services struct {
	Encryption EncryptionService
	FPE        FPEService
}

// EncryptionService defines the interface for encryption operations
type EncryptionService interface {
	// Symmetric encryption
	EncryptAES256GCM(plaintext, key, nonce []byte) ([]byte, error)
	DecryptAES256GCM(ciphertext, key, nonce []byte) ([]byte, error)
	
	// Stream cipher encryption
	EncryptChaCha20Poly1305(plaintext, key, nonce []byte) ([]byte, error)
	DecryptChaCha20Poly1305(ciphertext, key, nonce []byte) ([]byte, error)
	
	// Asymmetric encryption
	EncryptRSAOAEP(plaintext, publicKey []byte) ([]byte, error)
	DecryptRSAOAEP(ciphertext, privateKey []byte) ([]byte, error)
	
	// Quantum-resistant cryptography
	EncryptKyber(plaintext, publicKey []byte) ([]byte, error)
	DecryptKyber(ciphertext, privateKey []byte) ([]byte, error)
	
	// Key generation
	GenerateAESKey() ([]byte, error)
	GenerateRSAKeyPair() (privateKey, publicKey []byte, err error)
	GenerateKyberKeyPair() (privateKey, publicKey []byte, err error)
	
	// HSM integration
	EncryptWithHSM(plaintext []byte, keyID string) ([]byte, error)
	DecryptWithHSM(ciphertext []byte, keyID string) ([]byte, error)
	
	// Utility functions
	GenerateNonce(size int) ([]byte, error)
}

// FPEService defines the interface for format-preserving encryption operations
type FPEService interface {
	// Format-preserving encryption
	EncryptFPE(plaintext string, key []byte, tweak []byte) (string, error)
	DecryptFPE(ciphertext string, key []byte, tweak []byte) (string, error)
	
	// Common data type encryption
	EncryptCreditCard(cardNumber string, key []byte) (string, error)
	DecryptCreditCard(encryptedCard string, key []byte) (string, error)
	
	EncryptSSN(ssn string, key []byte) (string, error)
	DecryptSSN(encryptedSSN string, key []byte) (string, error)
}