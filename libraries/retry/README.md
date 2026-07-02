# MegaVerse Retry Library

Retry logic with exponential backoff.

## Features

- Exponential backoff
- Linear backoff
- Custom retry strategies
- Retry predicates
- Maximum attempts
- Jitter

## Usage

```typescript
import { retry } from '@megaverse/retry';

const result = await retry(
  () => unstableOperation(),
  {
    maxAttempts: 3,
    backoff: 'exponential',
    delay: 1000,
  }
);
```
