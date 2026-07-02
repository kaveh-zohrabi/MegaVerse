# MegaVerse Logging Configuration

## Overview

Centralized logging configuration for all services.

## Log Levels

- `debug`: Detailed debugging info
- `info`: General operational info
- `warn`: Warning messages
- `error`: Error messages
- `fatal`: Fatal errors

## Formats

### JSON (Production)
```json
{
  "timestamp": "2026-07-01T12:00:00Z",
  "level": "info",
  "service": "auth-service",
  "message": "User logged in",
  "userId": "123",
  "requestId": "abc-123"
}
```

### Text (Development)
```
2026-07-01 12:00:00 [INFO] auth-service: User logged in userId=123
```

## Destinations

- **Console**: Development
- **File**: Local debugging
- **Loki**: Centralized logging
- **Elasticsearch**: Search and analytics

## Configuration

```yaml
logging:
  level: info
  format: json
  outputs:
    - console
    - loki
  loki:
    url: http://loki:3100
```
