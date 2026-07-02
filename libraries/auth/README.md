# MegaVerse Auth Library

Authentication and authorization utilities.

## Features

- JWT handling
- OAuth 2.0
- Session management
- Token refresh
- Permission checking
- Role-based access

## Usage

```typescript
import { AuthService } from '@megaverse/auth';

const auth = new AuthService({
  jwtSecret: process.env.JWT_SECRET,
});

// Generate token
const token = auth.generateToken({ userId: '123', role: 'admin' });

// Verify token
const payload = auth.verifyToken(token);

// Check permission
const hasPermission = auth.hasPermission(payload, 'posts:write');
```
