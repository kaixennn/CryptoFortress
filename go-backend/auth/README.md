# Authentication Service

The Authentication Service is a core component of the CryptoFortress encryption suite, providing comprehensive authentication and authorization capabilities.

## Features

- JWT with refresh token rotation & automatic revocation
- OAuth2.0, SAML, and LDAP integration
- Multi-factor authentication (TOTP, WebAuthn)
- Role-based access control (RBAC) with fine-grained permissions

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Refresh access token
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/logout` - User logout

### Multi-Factor Authentication
- `POST /api/v1/auth/mfa/totp/enable` - Enable TOTP
- `POST /api/v1/auth/mfa/totp/verify` - Verify TOTP token
- `POST /api/v1/auth/mfa/totp/disable` - Disable TOTP
- `POST /api/v1/auth/mfa/webauthn/register` - Register WebAuthn credential
- `POST /api/v1/auth/mfa/webauthn/register/verify` - Verify WebAuthn registration
- `POST /api/v1/auth/mfa/webauthn/authenticate` - Authenticate with WebAuthn
- `POST /api/v1/auth/mfa/webauthn/authenticate/verify` - Verify WebAuthn authentication

### Role-Based Access Control
- `POST /api/v1/auth/rbac/roles` - Create role
- `DELETE /api/v1/auth/rbac/roles` - Delete role
- `POST /api/v1/auth/rbac/roles/assign` - Assign role to user
- `POST /api/v1/auth/rbac/roles/remove` - Remove role from user
- `POST /api/v1/auth/rbac/permissions` - Create permission
- `POST /api/v1/auth/rbac/permissions/assign` - Assign permission to role
- `POST /api/v1/auth/rbac/permissions/remove` - Remove permission from role
- `POST /api/v1/auth/rbac/permissions/check` - Check user permission
- `POST /api/v1/auth/rbac/users/roles` - Get user roles
- `POST /api/v1/auth/rbac/roles/permissions` - Get role permissions

## Environment Variables

- `AUTH_SERVICE_PORT` - Service port (default: 8080)
- `JWT_SECRET` - Secret key for JWT signing (required)
- `REFRESH_TOKEN_TTL` - Refresh token time-to-live in hours (default: 720)
- `ACCESS_TOKEN_TTL` - Access token time-to-live in minutes (default: 15)
- `DATABASE_URL` - Database connection URL
- `LDAP_SERVER` - LDAP server address
- `LDAP_PORT` - LDAP server port (default: 389)
- `OAUTH_CLIENT_ID` - OAuth2 client ID
- `OAUTH_CLIENT_SECRET` - OAuth2 client secret
- `SAML_ENTITY_ID` - SAML entity ID
- `VAULT_ADDR` - HashiCorp Vault address
- `VAULT_TOKEN` - HashiCorp Vault token

## Running the Service

### With Docker

```bash
docker build -t cryptofortress-auth .
docker run -p 8080:8080 cryptofortress-auth
```

### Locally

```bash
go run cmd/main.go
```