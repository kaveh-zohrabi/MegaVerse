# MegaVerse Cache Library

Multi-layer caching with Redis and in-memory support.

## Features

- Redis cache
- In-memory LRU cache
- Cache-aside pattern
- Cache invalidation
- TTL support
- Cache warming
- Distributed locks

## Usage

```typescript
import { CacheManager } from '@megaverse/cache';

const cache = new CacheManager({
  redis: { host: 'localhost', port: 6379 },
  memory: { maxSize: 1000 },
});

// Get or set
const user = await cache.getOrSet(`user:${id}`, () => fetchUser(id), {
  ttl: 3600,
});
```
