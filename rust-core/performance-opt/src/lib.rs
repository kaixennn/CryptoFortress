//! Performance Optimizations - Hardware acceleration and parallel processing
//!
//! This crate provides optimized implementations for bulk cryptographic
//! operations using hardware acceleration and parallel processing.

use crypto_engine::{bulk_encrypt, CryptoError};
use rayon::prelude::*;

/// Hardware-accelerated AES-NI implementation (mock)
#[cfg(target_arch = "x86_64")]
pub fn aes_ni_supported() -> bool {
    // In a real implementation, you would check for AES-NI support
    // using CPUID instruction
    true
}

/// Hardware-accelerated AES-NI implementation (mock)
#[cfg(not(target_arch = "x86_64"))]
pub fn aes_ni_supported() -> bool {
    false
}

/// Parallel bulk encryption using Rayon
pub fn parallel_bulk_encrypt(
    data_blocks: Vec<Vec<u8>>,
    key: &[u8],
) -> Result<Vec<Vec<u8>>, CryptoError> {
    // Use Rayon for parallel processing
    let results: Result<Vec<_>, _> = data_blocks
        .par_iter()
        .map(|block| {
            // In a real implementation, you would use hardware acceleration here
            let slice = block.as_slice();
            // For demonstration, we're just using the regular encryption
            // In practice, you would use specialized hardware-accelerated functions
            crypto_engine::aes_256_gcm_encrypt(slice, key, &[0u8; 12])
        })
        .collect();

    results
}

/// Memory-efficient bulk processing with streaming
pub struct StreamingProcessor {
    buffer: Vec<u8>,
    buffer_size: usize,
    key: Vec<u8>,
}

impl StreamingProcessor {
    pub fn new(key: &[u8], buffer_size: usize) -> Self {
        Self {
            buffer: Vec::with_capacity(buffer_size),
            buffer_size,
            key: key.to_vec(),
        }
    }

    pub fn process_chunk(&mut self, chunk: &[u8]) -> Result<Option<Vec<u8>>, CryptoError> {
        self.buffer.extend_from_slice(chunk);

        if self.buffer.len() >= self.buffer_size {
            // Process the buffer
            let data_to_process = self.buffer.drain(..self.buffer_size).collect::<Vec<_>>();
            let encrypted = crypto_engine::aes_256_gcm_encrypt(
                &data_to_process,
                &self.key,
                &[0u8; 12],
            )?;
            Ok(Some(encrypted))
        } else {
            Ok(None)
        }
    }

    pub fn finish(self) -> Result<Vec<u8>, CryptoError> {
        if self.buffer.is_empty() {
            Ok(Vec::new())
        } else {
            crypto_engine::aes_256_gcm_encrypt(
                &self.buffer,
                &self.key,
                &[0u8; 12],
            )
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parallel_bulk_encrypt() {
        let data_blocks = vec![
            b"Hello".to_vec(),
            b"World".to_vec(),
            b"Rust".to_vec(),
            b"Crypto".to_vec(),
        ];
        let key = [0u8; 32];

        let results = parallel_bulk_encrypt(data_blocks, &key).unwrap();
        assert_eq!(results.len(), 4);
    }

    #[test]
    fn test_streaming_processor() {
        let key = [0u8; 32];
        let mut processor = StreamingProcessor::new(&key, 10);

        let chunk1 = b"Hello";
        let chunk2 = b"World Rust Crypto";

        let result1 = processor.process_chunk(chunk1).unwrap();
        assert!(result1.is_none()); // Buffer not full yet

        let result2 = processor.process_chunk(chunk2).unwrap();
        assert!(result2.is_some()); // Buffer is now full

        let final_result = processor.finish().unwrap();
        assert!(!final_result.is_empty());
    }
}