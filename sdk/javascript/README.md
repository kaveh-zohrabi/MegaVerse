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
import { MegaVerseClient } from '@megaverse/javascript-sdk';

const client = new MegaVerseClient({
  baseUrl: 'http://localhost:8080',
});

// Login
await client.login('user@example.com', 'password');

// Get user
const user = await client.getUser('user-id');

// Create post
const post = await client.createPost('Hello MegaVerse!');

// Get feed
const { posts } = await client.getFeed({ limit: 10 });
```

## API Reference

### Client

```typescript
new MegaVerseClient({ baseUrl: string, apiKey?: string })
```

### Auth

- `register(email, name, password)` - Register new user
- `login(email, password)` - Login and get tokens
- `setToken(token)` - Set auth token

### Users

- `getUser(id)` - Get user profile
- `updateProfile(updates)` - Update own profile
- `follow(userId)` - Follow user
- `unfollow(userId)` - Unfollow user

### Posts

- `createPost(content, mediaUrl?)` - Create post
- `getPost(id)` - Get post
- `deletePost(id)` - Delete post
- `getFeed(params?)` - Get feed
- `getUserPosts(userId, params?)` - Get user posts

### Health

- `health()` - Check API health
