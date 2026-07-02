# Auth Service

Go service handling authentication, authorization, and JWT token management.

## Tech Stack

- **Language**: Go 1.22
- **Router**: gorilla/mux
- **Database**: MySQL 8.0
- **Auth**: JWT (HS256) + bcrypt

## Features

- User registration with email validation
- User login with password verification
- JWT access token generation
- JWT refresh token generation
- Token validation endpoint
- Protected route middleware

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | /auth/register | No | Register new user |
| POST | /auth/login | No | Login and get tokens |
| GET | /auth/validate | Yes | Validate JWT token |
| GET | /health | No | Health check |

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| DATABASE_HOST | localhost | MySQL host |
| DATABASE_PORT | 3306 | MySQL port |
| DATABASE_NAME | megaverse | Database name |
| DATABASE_USER | megaverse | Database user |
| DATABASE_PASSWORD | password | Database password |
| JWT_SECRET | dev-secret | JWT signing secret |
| PORT | 8081 | Server port |

## Development

```bash
# Run locally
go run ./cmd/server

# Build
go build -o bin/server ./cmd/server
```

## Database

Run migration before starting:

```bash
mysql -u root -p megaverse < migrations/001_create_users.sql
```
