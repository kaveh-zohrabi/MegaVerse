# Auth Service

Handles authentication, authorization, and session management.

## Features

- User registration and login
- OAuth 2.0 / OIDC support (Google, GitHub, etc.)
- JWT token management
- Session management
- MFA (Multi-Factor Authentication)
- Password reset flow
- API key management

## Endpoints

- `POST /auth/register` - Register new user
- `POST /auth/login` - User login
- `POST /auth/logout` - User logout
- `POST /auth/refresh` - Refresh token
- `POST /auth/forgot-password` - Password reset request
- `POST /auth/reset-password` - Password reset
- `POST /auth/mfa/enable` - Enable MFA
- `POST /auth/mfa/verify` - Verify MFA code
- `GET /auth/oauth/:provider` - OAuth login
- `GET /auth/oauth/:provider/callback` - OAuth callback

## Security

- Bcrypt password hashing
- JWT with RS256 signing
- Rate limiting on auth endpoints
- Account lockout after failed attempts
- CSRF protection
- Secure cookie handling
