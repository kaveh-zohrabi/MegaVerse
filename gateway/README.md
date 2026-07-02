# MegaVerse Gateway

API Gateway for routing, authentication, and rate limiting.

## Features

- Request routing
- Load balancing
- Authentication
- Rate limiting
- Request transformation
- Response caching
- Circuit breaking
- API versioning

## Architecture

```
Client → Gateway → Router → Service
              ↓
         Auth Plugin
              ↓
         Rate Limit Plugin
              ↓
         Transform Plugin
              ↓
         Forward
```

## Plugins

- JWT authentication
- OAuth 2.0
- API key validation
- Rate limiting (per-user, per-IP)
- Request/response transformation
- Response caching
- Circuit breaker

## Configuration

```yaml
routes:
  - path: /api/users
    service: user-service
    methods: [GET, POST, PUT, DELETE]
    plugins:
      - auth
      - rate-limit

  - path: /api/posts
    service: social-service
    methods: [GET, POST, PUT, DELETE]
    plugins:
      - auth
      - rate-limit
      - cache
```

## Development

```bash
pnpm dev
```

## Testing

```bash
pnpm test
```
