# MegaVerse JavaScript SDK

Official JavaScript/TypeScript SDK for the MegaVerse API.

## Installation

```bash
npm install @megaverse/javascript-sdk
# or
pnpm add @megaverse/javascript-sdk
```

## Quick Start

```typescript
import { MegaVerse } from '@megaverse/javascript-sdk';

const client = new MegaVerse({
  apiKey: 'your-api-key',
  baseUrl: 'https://api.megaverse.dev',
});

// Get user
const user = await client.users.get('user-id');

// Create post
const post = await client.posts.create({
  content: 'Hello MegaVerse!',
});
```

## Features

- TypeScript support
- Auto-retry with exponential backoff
- Request/response interceptors
- WebSocket support
- Type-safe API calls

## Documentation

See [docs.megaverse.dev/sdk/javascript](https://docs.megaverse.dev/sdk/javascript)
