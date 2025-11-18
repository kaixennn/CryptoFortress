# Key Management Service

The Key Management Service is a core component of the CryptoFortress encryption suite, providing comprehensive key management capabilities with advanced features for enterprise security.

## Features

- Automatic key rotation with versioning
- Shamir's Secret Sharing for key distribution
- Key expiration and revocation policies
- Cross-region key replication for disaster recovery

## API Endpoints

### Key Management
- `POST /api/v1/keymgmt/keys/generate` - Generate key
- `POST /api/v1/keymgmt/keys/generate-pair` - Generate key pair
- `POST /api/v1/keymgmt/keys/store` - Store key
- `POST /api/v1/keymgmt/keys/retrieve` - Retrieve key
- `POST /api/v1/keymgmt/keys/delete` - Delete key

### Key Rotation
- `POST /api/v1/keymgmt/rotation/rotate` - Rotate key
- `POST /api/v1/keymgmt/rotation/schedule` - Schedule rotation
- `POST /api/v1/keymgmt/rotation/cancel` - Cancel rotation
- `POST /api/v1/keymgmt/rotation/schedule/get` - Get rotation schedule
- `POST /api/v1/keymgmt/rotation/auto-enable` - Enable auto rotation
- `POST /api/v1/keymgmt/rotation/auto-disable` - Disable auto rotation

### Shamir's Secret Sharing
- `POST /api/v1/keymgmt/shamir/split` - Split secret
- `POST /api/v1/keymgmt/shamir/combine` - Combine shares
- `POST /api/v1/keymgmt/shamir/distribute` - Distribute key
- `POST /api/v1/keymgmt/shamir/recover` - Recover key

### Key Replication
- `POST /api/v1/keymgmt/replication/replicate` - Replicate key
- `POST /api/v1/keymgmt/replication/enable` - Enable cross-region replication
- `POST /api/v1/keymgmt/replication/disable` - Disable cross-region replication
- `POST /api/v1/keymgmt/replication/backup` - Backup key
- `POST /api/v1/keymgmt/replication/restore` - Restore key

## Environment Variables

- `KEYMGMT_SERVICE_PORT` - Service port (default: 8082)
- `DATABASE_URL` - Database connection URL
- `VAULT_ADDR` - HashiCorp Vault address
- `VAULT_TOKEN` - HashiCorp Vault token
- `KEY_ROTATION_PERIOD` - Key rotation period in days (default: 90)
- `SHAMIR_THRESHOLD` - Shamir threshold (default: 2)
- `SHAMIR_SHARES` - Shamir shares (default: 3)
- `REPLICATION_ENABLED` - Enable replication (default: false)
- `REPLICATION_REGIONS` - Comma-separated list of replication regions

## Running the Service

### With Docker

```bash
docker build -t cryptofortress-keymgmt .
docker run -p 8082:8082 cryptofortress-keymgmt
```

### Locally

```bash
go run cmd/main.go
```