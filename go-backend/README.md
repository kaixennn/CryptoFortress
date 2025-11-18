# CryptoFortress Go Backend Services

This directory contains the Go microservices that form the backend of the CryptoFortress encryption suite.

## Services

1. **Authentication Service** (`auth/`) - JWT with refresh token rotation, OAuth2.0, SAML, LDAP integration, MFA, RBAC
2. **Encryption Service** (`encryption/`) - Multiple algorithms, quantum-resistant cryptography, FPE, HSM integration
3. **Key Management Service** (`keymgmt/`) - Key rotation, Shamir's Secret Sharing, key expiration, cross-region replication
4. **Audit & Compliance Service** (`audit/`) - Immutable audit trails, compliance reporting, SIEM integration

## Prerequisites

- Go 1.21 or later
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (for database storage)

## Running Services Locally

### Option 1: Using Docker Compose (Recommended)

```bash
# Navigate to the go-backend directory
cd go-backend

# Build and start all services
docker-compose up --build

# To run in detached mode
docker-compose up --build -d

# To stop all services
docker-compose down
```

### Option 2: Running Services Individually

```bash
# Navigate to each service directory and run
cd auth
go run cmd/main.go

cd encryption
go run cmd/main.go

cd keymgmt
go run cmd/main.go

cd audit
go run cmd/main.go
```

## Service Ports

- Authentication Service: http://localhost:8080
- Encryption Service: http://localhost:8081
- Key Management Service: http://localhost:8082
- Audit & Compliance Service: http://localhost:8083
- PostgreSQL Database: postgres://localhost:5432
- pgAdmin: http://localhost:5050

## Database Access

The services use PostgreSQL as their database. You can access the database using:

- **pgAdmin**: Visit http://localhost:5050 and log in with:
  - Email: admin@example.com
  - Password: admin
- **Direct connection**: 
  - Host: localhost:5432
  - Username: user
  - Password: password
  - Database: cryptofortress

## Health Checks

Each service exposes a health check endpoint at `/health`:

```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
```

## API Documentation

Each service has its own API documentation in its respective README.md file:

- [Authentication Service API](auth/README.md)
- [Encryption Service API](encryption/README.md)
- [Key Management Service API](keymgmt/README.md)
- [Audit & Compliance Service API](audit/README.md)

## Development

### Adding New Dependencies

```bash
# Navigate to the service directory
cd auth  # or encryption, keymgmt, audit

# Add new dependency
go get github.com/some/dependency

# Update go.mod and go.sum
go mod tidy
```

### Testing

```bash
# Run tests for a specific service
cd auth
go test ./...

# Run tests with coverage
go test ./... -cover
```

## Environment Variables

Each service uses environment variables for configuration. Check each service's README for specific variables.

Common variables across services:
- `SERVICE_PORT` - Port the service listens on
- `DATABASE_URL` - Database connection string
- `VAULT_ADDR` - HashiCorp Vault address
- `VAULT_TOKEN` - HashiCorp Vault token

## Troubleshooting

1. **Port conflicts**: Make sure ports 8080-8083 and 5432 are available
2. **Database connection issues**: Verify PostgreSQL is running and credentials are correct
3. **Service won't start**: Check logs with `docker-compose logs <service-name>`
4. **Dependency issues**: Run `go mod tidy` in the service directory

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests
5. Submit a pull request