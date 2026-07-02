# API Design

## REST API

### Base URL

```
https://api.megaverse.dev/v1
```

### Authentication

All API requests require authentication via:
- JWT token (Bearer)
- API key
- OAuth 2.0

### Response Format

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

### Error Format

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid email format",
    "details": { ... }
  }
}
```

### Pagination

```
GET /api/users?page=1&limit=20&sort=created_at&order=desc
```

### Filtering

```
GET /api/posts?status=published&author_id=123&tag=javascript
```

## GraphQL

### Endpoint

```
https://api.megaverse.dev/graphql
```

### Schema

```graphql
type User {
  id: ID!
  email: String!
  name: String!
  posts: [Post!]!
}

type Query {
  user(id: ID!): User
  users(limit: Int, offset: Int): [User!]!
}

type Mutation {
  createUser(input: CreateUserInput!): User!
  updateUser(id: ID!, input: UpdateUserInput!): User!
}
```

## gRPC

### Service Definitions

```protobuf
service UserService {
  rpc GetUser(GetUserRequest) returns (User);
  rpc CreateUser(CreateUserRequest) returns (User);
  rpc UpdateUser(UpdateUserRequest) returns (User);
  rpc DeleteUser(DeleteUserRequest) returns (google.protobuf.Empty);
}
```

## WebSocket

### Connection

```
wss://api.megaverse.dev/ws?token=xxx
```

### Events

```json
{
  "type": "message",
  "data": {
    "conversationId": "123",
    "content": "Hello!",
    "sender": { ... }
  }
}
```
