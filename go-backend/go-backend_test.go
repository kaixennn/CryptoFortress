package main

import (
	"testing"
)

// TestRustCryptoEngine tests the Rust FFI integration
func TestRustCryptoEngine(t *testing.T) {
	t.Run("EncryptDecryptRoundTrip", func(t *testing.T) {
		// Test normal encryption/decryption round trip
		engine := &RustCryptoEngine{}
		plaintext := []byte("Hello, CryptoFortress!")
		key := make([]byte, 32) // 256-bit key
		nonce := make([]byte, 12) // 96-bit nonce

		encrypted, err := engine.Encrypt(plaintext, key, nonce)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		if len(encrypted) == 0 {
			t.Error("Encrypted data is empty")
		}

		decrypted, err := engine.Decrypt(encrypted, key)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != string(plaintext) {
			t.Errorf("Decrypted data doesn't match original. Got %s, want %s", 
				string(decrypted), string(plaintext))
		}
	})

	t.Run("EncryptEmptyData", func(t *testing.T) {
		// Test encryption of empty data
		engine := &RustCryptoEngine{}
		plaintext := []byte("")
		key := make([]byte, 32)
		nonce := make([]byte, 12)

		_, err := engine.Encrypt(plaintext, key, nonce)
		if err == nil {
			t.Error("Expected error when encrypting empty data, but got none")
		}
	})

	t.Run("InvalidKeyLength", func(t *testing.T) {
		// Test with invalid key length
		engine := &RustCryptoEngine{}
		plaintext := []byte("test data")
		key := make([]byte, 16) // 128-bit key (invalid)
		nonce := make([]byte, 12)

		_, err := engine.Encrypt(plaintext, key, nonce)
		if err == nil {
			t.Error("Expected error with invalid key length, but got none")
		}
	})

	t.Run("InvalidNonceLength", func(t *testing.T) {
		// Test with invalid nonce length
		engine := &RustCryptoEngine{}
		plaintext := []byte("test data")
		key := make([]byte, 32)
		nonce := make([]byte, 8) // 64-bit nonce (invalid)

		_, err := engine.Encrypt(plaintext, key, nonce)
		if err == nil {
			t.Error("Expected error with invalid nonce length, but got none")
		}
	})
}

// TestAIServiceClient tests the Python AI service client
func TestAIServiceClient(t *testing.T) {
	t.Run("PatternDetection", func(t *testing.T) {
		// Test pattern detection functionality
		client := &AIServiceClient{
			BaseURL: "http://localhost:5000",
		}

		// This would normally make an HTTP request, but we're testing the structure
		data := map[string]interface{}{
			"algorithm": "AES-256-GCM",
			"data_size": 1024,
		}

		// In a real test, we would mock the HTTP response
		// For now, we're just testing that the method exists and has the right signature
		_ = client.DetectEncryptionPatterns(data)
	})

	t.Run("KeyStrengthAnalysis", func(t *testing.T) {
		// Test key strength analysis functionality
		client := &AIServiceClient{
			BaseURL: "http://localhost:5000",
		}

		keyData := map[string]interface{}{
			"key_hex": "0123456789abcdef0123456789abcdef",
			"algorithm": "AES-256",
		}

		_ = client.AnalyzeKeyStrength(keyData)
	})
}

// Mock AIServiceClient for testing
type AIServiceClient struct {
	BaseURL string
}

func (c *AIServiceClient) DetectEncryptionPatterns(data map[string]interface{}) error {
	// Mock implementation
	return nil
}

func (c *AIServiceClient) AnalyzeKeyStrength(keyData map[string]interface{}) error {
	// Mock implementation
	return nil
}