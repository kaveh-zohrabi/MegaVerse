# API Gateway

Go service serving as the single entry point for all client requests.

## Tech Stack

- **Language**: Go 1.22
- **Router**: gorilla/mux
- **Proxy**: net/http/httputil

## Features

- Request routing to downstream services
- Reverse proxy with header forwarding
- Rate limiting (per-IP)
- CORS handling
- Request logging
- Health check aggregation

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /health | Gateway + service health |
| * | /auth/* | Proxy to auth-service |
| * | /users/* | Proxy to user-service |
| * | /posts/* | Proxy to social-service |
| * | /feed/* | Proxy to social-service |

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| PORT | 8080 | Server port |
| AUTH_SERVICE_URL | http://localhost:8081 | Auth service URL |
| USER_SERVICE_URL | http://localhost:8082 | User service URL |
| RATE_LIMIT | 1000 | Max requests per window |
| RATE_WINDOW | 60 | Rate window in seconds |

## Architecture

```
Client → API Gateway → Auth Service (port 8081)
                    → User Service (port 8082)
                    → Social Service (port 8083)
```

## Development

```bash
go run ./cmd/server
```
