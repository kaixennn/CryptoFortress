# Encryption Service

The Encryption Service is a core component of the CryptoFortress encryption suite, providing comprehensive encryption capabilities with support for multiple algorithms and advanced features.

## Features

- Multiple algorithms: AES-256-GCM, ChaCha20-Poly1305, RSA-OAEP
- Quantum-resistant cryptography (Kyber, Dilithium)
- Format-preserving encryption for databases
- Hardware Security Module (HSM) integration

## API Endpoints

### Basic Encryption
- `POST /api/v1/encryption/encrypt` - Encrypt data
- `POST /api/v1/encryption/decrypt` - Decrypt data
- `POST /api/v1/encryption/generate-key` - Generate encryption key

### Format-Preserving Encryption
- `POST /api/v1/encryption/fpe/encrypt` - FPE encryption
- `POST /api/v1/encryption/fpe/decrypt` - FPE decryption

### Specialized Encryption
- `POST /api/v1/encryption/credit-card/encrypt` - Encrypt credit card
- `POST /api/v1/encryption/credit-card/decrypt` - Decrypt credit card
- `POST /api/v1/encryption/ssn/encrypt` - Encrypt SSN
- `POST /api/v1/encryption/ssn/decrypt` - Decrypt SSN

## Environment Variables

- `ENCRYPTION_SERVICE_PORT` - Service port (default: 8081)
- `DATABASE_URL` - Database connection URL
- `HSM_ENABLED` - Enable HSM integration (default: false)
- `HSM_ADDRESS` - HSM address
- `HSM_TOKEN` - HSM authentication token
- `VAULT_ADDR` - HashiCorp Vault address
- `VAULT_TOKEN` - HashiCorp Vault token
- `AES_KEY_SIZE` - AES key size in bits (default: 256)
- `RSA_KEY_SIZE` - RSA key size in bits (default: 2048)
- `DEFAULT_ALGORITHM` - Default encryption algorithm (default: AES-256-GCM)

## Running the Service

### With Docker

```bash
docker build -t cryptofortress-encryption .
docker run -p 8081:8081 cryptofortress-encryption
```

### Locally

```bash
go run cmd/main.go
```