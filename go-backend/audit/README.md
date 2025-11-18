# Audit & Compliance Service

The Audit & Compliance Service is a core component of the CryptoFortress encryption suite, providing comprehensive audit trail management and compliance reporting capabilities.

## Features

- Immutable audit trails with cryptographic hashing
- GDPR, HIPAA, SOC2 compliance reporting
- Real-time security event monitoring
- SIEM integration (Splunk, Elasticsearch)

## API Endpoints

### Audit Management
- `POST /api/v1/audit/events/log` - Log audit event
- `POST /api/v1/audit/events/trail` - Get audit trail
- `POST /api/v1/audit/events/immutable` - Create immutable trail
- `POST /api/v1/audit/events/verify` - Verify trail integrity

### Compliance Management
- `POST /api/v1/audit/compliance/report` - Generate compliance report
- `POST /api/v1/audit/compliance/status` - Get compliance status
- `GET /api/v1/audit/compliance/standards` - List compliance standards
- `POST /api/v1/audit/compliance/gdpr/request` - Handle data subject request
- `POST /api/v1/audit/compliance/gdpr/inventory` - Generate data inventory

### SIEM Integration
- `POST /api/v1/audit/siem/event` - Send SIEM event
- `GET /api/v1/audit/siem/threats` - Monitor threats
- `POST /api/v1/audit/siem/endpoint` - Configure SIEM endpoint
- `POST /api/v1/audit/siem/alert-rule` - Create alert rule

## Environment Variables

- `AUDIT_SERVICE_PORT` - Service port (default: 8083)
- `DATABASE_URL` - Database connection URL
- `AUDIT_LOG_RETENTION` - Audit log retention in days (default: 365)
- `COMPLIANCE_STANDARDS` - Comma-separated list of compliance standards (default: GDPR,HIPAA,SOC2)
- `SIEM_ENABLED` - Enable SIEM integration (default: false)
- `SIEM_ENDPOINTS` - Comma-separated list of SIEM endpoints
- `IMMUTABLE_AUDIT_TRAILS` - Enable immutable audit trails (default: true)

## Running the Service

### With Docker

```bash
docker build -t cryptofortress-audit .
docker run -p 8083:8083 cryptofortress-audit
```

### Locally

```bash
go run cmd/main.go
```