# Messaging Service

Go service handling real-time messaging with WebSocket support.

## Tech Stack

- **Language**: Go 1.22
- **Router**: gorilla/mux
- **WebSocket**: gorilla/websocket
- **Database**: MySQL 8.0

## Features

- Direct and group conversations
- Real-time messaging via WebSocket
- Message history with pagination
- Read receipts
- Room-based broadcasting

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /health | No | Health check |
| WS | /ws | No | WebSocket connection |
| POST | /conversations | Yes | Create conversation |
| GET | /conversations | Yes | List conversations |
| POST | /conversations/{id}/messages | Yes | Send message |
| GET | /conversations/{id}/messages | Yes | Get messages |

## WebSocket

Connect to `ws://localhost:8084/ws?user_id=xxx&room=xxx`

### Message Format

```json
{
  "type": "message",
  "room": "global",
  "data": {
    "user_id": "xxx",
    "content": "Hello!"
  }
}
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| DATABASE_HOST | localhost | MySQL host |
| PORT | 8084 | Server port |

## Development

```bash
go run ./cmd/server
```
