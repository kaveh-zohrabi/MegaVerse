# Messaging Service

Real-time messaging between users.

## Features

- Direct messages
- Group chats
- Message history
- Read receipts
- Typing indicators
- File/image sharing
- Message reactions
- Threaded conversations

## Endpoints

- `GET /conversations` - List conversations
- `POST /conversations` - Create conversation
- `GET /conversations/:id/messages` - Get messages
- `POST /conversations/:id/messages` - Send message
- `PUT /messages/:id` - Edit message
- `DELETE /messages/:id` - Delete message
- `POST /messages/:id/reactions` - Add reaction
- `WebSocket /ws` - Real-time connection
