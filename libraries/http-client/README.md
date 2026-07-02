# MegaVerse HTTP Client

HTTP client with retry, circuit breaker, and logging.

## Features

- Request/response interceptors
- Automatic retries
- Circuit breaker
- Timeout handling
- Request/response logging
- Rate limiting
- Connection pooling

## Usage

```typescript
import { HttpClient } from '@megaverse/http-client';

const client = new HttpClient({
  baseUrl: 'https://api.example.com',
  retries: 3,
  timeout: 5000,
});

const response = await client.get('/users/123');
```
