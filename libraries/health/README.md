# MegaVerse Health

Health check endpoints and utilities.

## Features

- Liveness checks
- Readiness checks
- Dependency checks
- Custom health checks
- Health aggregation

## Usage

```typescript
import { HealthCheck } from '@megaverse/health';

const health = new HealthCheck();

// Register checks
health.register('database', async () => {
  await db.ping();
});

health.register('redis', async () => {
  await redis.ping();
});

// Get health status
const status = await health.check();
// { status: 'healthy', checks: { database: 'ok', redis: 'ok' } }
```
