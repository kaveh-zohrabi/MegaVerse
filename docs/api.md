# MegaVerse API Reference

## Base URL

```
https://api.megaverse.dev/v1
```

## Authentication

All requests require authentication via Bearer token:

```
Authorization: Bearer <token>
```

## Endpoints

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /users/:id | Get user profile |
| PUT | /users/:id | Update profile |
| DELETE | /users/:id | Delete account |

### Posts

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /posts | List posts |
| POST | /posts | Create post |
| GET | /posts/:id | Get post |
| PUT | /posts/:id | Update post |
| DELETE | /posts/:id | Delete post |

### Messages

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /conversations | List conversations |
| POST | /conversations | Create conversation |
| GET | /conversations/:id/messages | Get messages |
| POST | /conversations/:id/messages | Send message |

### AI

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /ai/embeddings | Generate embeddings |
| POST | /ai/search | Semantic search |
| POST | /ai/recommend | Get recommendations |

## Response Format

```json
{
  "success": true,
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100
  }
}
```

## Error Format

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input"
  }
}
```

## Rate Limiting

- 1000 requests per minute per API key
- Headers: `X-RateLimit-Limit`, `X-RateLimit-Remaining`
