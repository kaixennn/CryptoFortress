package main

/*
#cgo LDFLAGS: -L../rust-core/target/release -lffi_bridge
#include "../rust-core/ffi-bridge.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// RustCryptoEngine wraps the Rust crypto functions
type RustCryptoEngine struct{}

// Encrypt uses Rust AES-256-GCM encryption
func (r *RustCryptoEngine) Encrypt(plaintext, key, nonce []byte) ([]byte, error) {
	// Convert Go slices to C pointers
	cPlaintext := C.CBytes(plaintext)
	cKey := C.CBytes(key)
	cNonce := C.CBytes(nonce)
	defer C.free(unsafe.Pointer(cPlaintext))
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cNonce))

	// Call Rust encryption function
	result := C.crypto_aes_256_gcm_encrypt(
		(*C.uchar)(cPlaintext),
		C.size_t(len(plaintext)),
		(*C.uchar)(cKey),
		C.size_t(len(key)),
		(*C.uchar)(cNonce),
		C.size_t(len(nonce)),
	)

	// Handle result
	if !result.success {
		return nil, fmt.Errorf("encryption failed with error code: %d", result.error_code)
	}

	// Convert result back to Go bytes
	encrypted := C.GoBytes(unsafe.Pointer(result.data), C.int(result.len))

	// Clean up memory allocated by Rust
	C.crypto_free(result.data, result.len)

	return encrypted, nil
}

// Decrypt uses Rust AES-256-GCM decryption
func (r *RustCryptoEngine) Decrypt(ciphertext, key []byte) ([]byte, error) {
	// Convert Go slices to C pointers
	cCiphertext := C.CBytes(ciphertext)
	cKey := C.CBytes(key)
	defer C.free(unsafe.Pointer(cCiphertext))
	defer C.free(unsafe.Pointer(cKey))

	// Call Rust decryption function
	result := C.crypto_aes_256_gcm_decrypt(
		(*C.uchar)(cCiphertext),
		C.size_t(len(ciphertext)),
		(*C.uchar)(cKey),
		C.size_t(len(key)),
	)

	// Handle result
	if !result.success {
		return nil, fmt.Errorf("decryption failed with error code: %d", result.error_code)
	}

	// Convert result back to Go bytes
	decrypted := C.GoBytes(unsafe.Pointer(result.data), C.int(result.len))

	// Clean up memory allocated by Rust
	C.crypto_free(result.data, result.len)

	return decrypted, nil
}

func main() {
	// Example usage
	engine := &RustCryptoEngine{}

	// Sample data
	plaintext := []byte("Hello, CryptoFortress!")
	key := make([]byte, 32) // 256-bit key
	nonce := make([]byte, 12) // 96-bit nonce

	// Encrypt
	encrypted, err := engine.Encrypt(plaintext, key, nonce)
	if err != nil {
		fmt.Printf("Encryption error: %v\n", err)
		return
	}

	fmt.Printf("Encrypted %d bytes\n", len(encrypted))

	// Decrypt
	decrypted, err := engine.Decrypt(encrypted, key)
	if err != nil {
		fmt.Printf("Decryption error: %v\n", err)
		return
	}

	fmt.Printf("Decrypted: %s\n", string(decrypted))
}