# MegaVerse Circuit Breaker

Circuit breaker pattern implementation.

## Features

- State management (closed, open, half-open)
- Automatic recovery
- Configurable thresholds
- Fallback support
- Metrics tracking

## Usage

```typescript
import { CircuitBreaker } from '@megaverse/circuit-breaker';

const breaker = new CircuitBreaker({
  failureThreshold: 5,
  resetTimeout: 30000,
});

const result = await breaker.execute(() => externalService.call());
```
