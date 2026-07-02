# User Service

Go service managing user profiles, preferences, and social connections.

## Tech Stack

- **Language**: Go 1.22
- **Router**: gorilla/mux
- **Database**: MySQL 8.0

## Features

- User profile management
- Profile updates (bio, website, location, avatar)
- Follow/unfollow users
- Follower/following lists
- Automatic profile creation

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /users/{id} | No | Get user profile |
| PUT | /users/me | Yes | Update own profile |
| POST | /users/{id}/follow | Yes | Follow user |
| POST | /users/{id}/unfollow | Yes | Unfollow user |
| GET | /users/{id}/followers | No | Get followers list |
| GET | /users/{id}/following | No | Get following list |
| GET | /health | No | Health check |

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| DATABASE_HOST | localhost | MySQL host |
| DATABASE_PORT | 3306 | MySQL port |
| DATABASE_NAME | megaverse | Database name |
| AUTH_SERVICE_URL | http://localhost:8081 | Auth service URL |
| PORT | 8082 | Server port |

## Development

```bash
go run ./cmd/server
```
