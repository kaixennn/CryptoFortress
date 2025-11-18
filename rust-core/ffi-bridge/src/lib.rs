//! FFI Bridge Layer - C ABI compatibility for CryptoFortress
//!
//! This crate provides C-compatible interfaces for the crypto-engine
//! to enable integration with Go, Python, and Node.js.

use crypto_engine::{
    aes_256_gcm_encrypt, aes_256_gcm_decrypt, 
    constant_time_compare, secure_zero_memory,
    generate_random_bytes, CryptoError
};
use std::ffi::{CStr, CString};
use std::os::raw::{c_char, c_int, c_void};

/// C-compatible result type
#[repr(C)]
pub struct CryptoResult {
    pub success: bool,
    pub data: *mut u8,
    pub len: usize,
    pub error_code: i32,
}

/// Free memory allocated by Rust
#[no_mangle]
pub extern "C" fn crypto_free(ptr: *mut u8, len: usize) {
    if !ptr.is_null() {
        unsafe {
            let _ = Vec::from_raw_parts(ptr, 0, len);
        }
    }
}

/// AES-256-GCM encryption
///
/// # Safety
/// - plaintext, key, and nonce must be valid pointers
/// - plaintext_len, key_len, and nonce_len must be correct
#[no_mangle]
pub unsafe extern "C" fn crypto_aes_256_gcm_encrypt(
    plaintext: *const u8,
    plaintext_len: usize,
    key: *const u8,
    key_len: usize,
    nonce: *const u8,
    nonce_len: usize,
) -> CryptoResult {
    // Validate inputs
    if plaintext.is_null() || key.is_null() || nonce.is_null() {
        return CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: 1, // Invalid input
        };
    }

    // Convert to slices
    let plaintext_slice = std::slice::from_raw_parts(plaintext, plaintext_len);
    let key_slice = std::slice::from_raw_parts(key, key_len);
    let nonce_slice = std::slice::from_raw_parts(nonce, nonce_len);

    // Perform encryption
    match aes_256_gcm_encrypt(plaintext_slice, key_slice, nonce_slice) {
        Ok(encrypted) => {
            let len = encrypted.len();
            let ptr = encrypted.as_ptr() as *mut u8;
            std::mem::forget(encrypted); // Transfer ownership to C

            CryptoResult {
                success: true,
                data: ptr,
                len,
                error_code: 0,
            }
        }
        Err(e) => CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: match e {
                CryptoError::InvalidKeyLength => 2,
                CryptoError::InvalidNonceLength => 3,
                CryptoError::EmptyPlaintext => 4,
                _ => 5,
            },
        },
    }
}

/// AES-256-GCM decryption
///
/// # Safety
/// - ciphertext and key must be valid pointers
/// - ciphertext_len and key_len must be correct
#[no_mangle]
pub unsafe extern "C" fn crypto_aes_256_gcm_decrypt(
    ciphertext: *const u8,
    ciphertext_len: usize,
    key: *const u8,
    key_len: usize,
) -> CryptoResult {
    // Validate inputs
    if ciphertext.is_null() || key.is_null() {
        return CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: 1, // Invalid input
        };
    }

    // Convert to slices
    let ciphertext_slice = std::slice::from_raw_parts(ciphertext, ciphertext_len);
    let key_slice = std::slice::from_raw_parts(key, key_len);

    // Perform decryption
    match aes_256_gcm_decrypt(ciphertext_slice, key_slice) {
        Ok(decrypted) => {
            let len = decrypted.len();
            let ptr = decrypted.as_ptr() as *mut u8;
            std::mem::forget(decrypted); // Transfer ownership to C

            CryptoResult {
                success: true,
                data: ptr,
                len,
                error_code: 0,
            }
        }
        Err(e) => CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: match e {
                CryptoError::InvalidKeyLength => 2,
                CryptoError::InvalidCiphertext => 6,
                CryptoError::EmptyPlaintext => 4,
                _ => 5,
            },
        },
    }
}

/// Constant-time comparison
///
/// # Safety
/// - a and b must be valid pointers
/// - a_len and b_len must be correct
#[no_mangle]
pub unsafe extern "C" fn crypto_constant_time_compare(
    a: *const u8,
    a_len: usize,
    b: *const u8,
    b_len: usize,
) -> bool {
    // Validate inputs
    if a.is_null() || b.is_null() {
        return false;
    }

    // Convert to slices
    let a_slice = std::slice::from_raw_parts(a, a_len);
    let b_slice = std::slice::from_raw_parts(b, b_len);

    constant_time_compare(a_slice, b_slice)
}

/// Generate random bytes
///
/// # Safety
/// - len must be reasonable (<= 1024)
#[no_mangle]
pub unsafe extern "C" fn crypto_generate_random_bytes(len: usize) -> CryptoResult {
    // Validate input
    if len > 1024 {
        return CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: 7, // Invalid length
        };
    }

    // Generate random bytes
    match generate_random_bytes(len) {
        Ok(bytes) => {
            let len = bytes.len();
            let ptr = bytes.as_ptr() as *mut u8;
            std::mem::forget(bytes); // Transfer ownership to C

            CryptoResult {
                success: true,
                data: ptr,
                len,
                error_code: 0,
            }
        }
        Err(_) => CryptoResult {
            success: false,
            data: std::ptr::null_mut(),
            len: 0,
            error_code: 5, // General error
        },
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ffi_encryption() {
        let plaintext = b"Hello, World!";
        let key = [0u8; 32];
        let nonce = [0u8; 12];

        let result = unsafe {
            crypto_aes_256_gcm_encrypt(
                plaintext.as_ptr(),
                plaintext.len(),
                key.as_ptr(),
                key.len(),
                nonce.as_ptr(),
                nonce.len(),
            )
        };

        assert!(result.success);
        assert!(result.len > 0);
        assert_eq!(result.error_code, 0);

        // Clean up
        unsafe {
            crypto_free(result.data, result.len);
        }
    }

    #[test]
    fn test_ffi_decryption() {
        let ciphertext = b"encrypted_data_mock";
        let key = [0u8; 32];

        let result = unsafe {
            crypto_aes_256_gcm_decrypt(
                ciphertext.as_ptr(),
                ciphertext.len(),
                key.as_ptr(),
                key.len(),
            )
        };

        // In our mock implementation, this will succeed
        assert!(result.success);
        assert_eq!(result.error_code, 0);

        // Clean up
        unsafe {
            crypto_free(result.data, result.len);
        }
    }
}