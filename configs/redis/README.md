# Redis Configuration

## Overview

Redis configuration for caching and sessions.

## Configuration

```redis
# Network
bind 0.0.0.0
port 6379
tcp-backlog 511
timeout 0
tcp-keepalive 300

# Memory
maxmemory 2gb
maxmemory-policy allkeys-lru

# Persistence
save 900 1
save 300 10
save 60 10000

# Security
requirepass ${REDIS_PASSWORD}
```

## Usage

### Session Store
```typescript
import Redis from 'ioredis';

const redis = new Redis({
  host: 'localhost',
  port: 6379,
  password: process.env.REDIS_PASSWORD,
});
```

### Cache
```typescript
await redis.set('user:123', JSON.stringify(userData), 'EX', 3600);
const user = JSON.parse(await redis.get('user:123'));
```
