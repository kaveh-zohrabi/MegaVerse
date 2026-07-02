# API Gateway

The API Gateway serves as the single entry point for all client requests.

## Features

- Request routing and load balancing
- Authentication and authorization
- Rate limiting
- Request/response transformation
- Circuit breaking
- API versioning
- CORS handling

## Architecture

```
Client → API Gateway → Service Router → Backend Service
              ↓
         Auth Check
              ↓
         Rate Limit
              ↓
         Transform
              ↓
         Forward
```

## Endpoints

- `GET /health` - Health check
- `GET /ready` - Readiness check
- All API routes are proxied to appropriate services

## Configuration

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| GATEWAY_PORT | Server port | 8080 |
| GATEWAY_RATE_LIMIT | Max requests per window | 1000 |
| GATEWAY_RATE_WINDOW | Rate window in seconds | 60 |
| DATABASE_URL | PostgreSQL connection | - |
| REDIS_URL | Redis connection | - |

## Development

```bash
go run ./cmd/server
```

## Testing

```bash
go test ./...
```
