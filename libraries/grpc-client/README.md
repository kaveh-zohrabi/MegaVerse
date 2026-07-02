# MegaVerse gRPC Client

gRPC client with automatic retries and connection management.

## Features

- Protocol Buffers support
- Automatic retries
- Load balancing
- Connection pooling
- Streaming support
- Deadline propagation

## Usage

```typescript
import { GrpcClient } from '@megaverse/grpc-client';

const client = new GrpcClient({
  service: 'user-service',
  proto: './protos/user.proto',
});

const user = await client.call('GetUser', { id: '123' });
```
