# MegaVerse Rate Limiter

Distributed rate limiting with Redis.

## Features

- Token bucket algorithm
- Sliding window
- Per-user/per-IP limiting
- Distributed rate limiting
- Custom rules
- Rate limit headers

## Usage

```typescript
import { RateLimiter } from '@megaverse/rate-limiter';

const limiter = new RateLimiter({
  redis: { host: 'localhost' },
  max: 100,
  window: 60, // seconds
});

// In middleware
app.use(limiter.middleware());
```
