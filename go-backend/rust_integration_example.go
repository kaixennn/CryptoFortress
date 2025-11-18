package main

/*
#cgo LDFLAGS: -L./rust-core/target/release -lffi_bridge
#include "./rust-core/ffi-bridge.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Example of calling Rust crypto functions from Go
	plaintext := []byte("Hello, CryptoFortress!")
	key := make([]byte, 32) // 256-bit key
	nonce := make([]byte, 12) // 96-bit nonce

	// Call Rust AES-256-GCM encryption
	cPlaintext := C.CBytes(plaintext)
	cKey := C.CBytes(key)
	cNonce := C.CBytes(nonce)

	result := C.crypto_aes_256_gcm_encrypt(
		(*C.uchar)(cPlaintext),
		C.size_t(len(plaintext)),
		(*C.uchar)(cKey),
		C.size_t(len(key)),
		(*C.uchar)(cNonce),
		C.size_t(len(nonce)),
	)

	if result.success {
		// Convert result back to Go bytes
		encrypted := C.GoBytes(unsafe.Pointer(result.data), C.int(result.len))
		fmt.Printf("Encrypted %d bytes\n", len(encrypted))

		// Clean up memory allocated by Rust
		C.crypto_free(result.data, result.len)
	} else {
		fmt.Printf("Encryption failed with error code: %d\n", result.error_code)
	}

	// Clean up C memory
	C.free(unsafe.Pointer(cPlaintext))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cNonce))
}