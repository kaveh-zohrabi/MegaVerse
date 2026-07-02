# MegaVerse Middleware

Express/Fastify middleware collection.

## Features

- Authentication middleware
- Rate limiting
- Request validation
- CORS
- Helmet security
- Compression
- Request logging
- Error handling

## Usage

```typescript
import { 
  authMiddleware,
  rateLimitMiddleware,
  corsMiddleware,
  errorHandler
} from '@megaverse/middleware';

app.use(corsMiddleware());
app.use(rateLimitMiddleware({ max: 100 }));
app.use('/api', authMiddleware);
app.use(errorHandler);
```
