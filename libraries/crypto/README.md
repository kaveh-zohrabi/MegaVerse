# MegaVerse Crypto Library

Cryptographic operations and utilities.

## Features

- Password hashing (bcrypt, argon2)
- Token generation
- Encryption/decryption
- Hash functions
- HMAC signing
- Key derivation

## Usage

```typescript
import { hash, verify, generateToken } from '@megaverse/crypto';

// Password hashing
const hashed = await hash('password123');
const isValid = await verify('password123', hashed);

// Token generation
const token = generateToken({ length: 32 });
```
