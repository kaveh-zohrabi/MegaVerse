# MegaVerse Logger

Structured logging library with multiple transports.

## Features

- Multiple log levels
- JSON structured output
- Console/file transports
- Request correlation IDs
- Performance logging
- Error context enrichment

## Usage

```typescript
import { createLogger } from '@megaverse/logger';

const logger = createLogger({
  service: 'user-service',
  level: 'info',
});

logger.info('User created', { userId: '123' });
logger.error('Failed to process', { error, requestId });
```
