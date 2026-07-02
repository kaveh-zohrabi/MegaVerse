# MegaVerse Queue Library

Message queue abstraction for Kafka, RabbitMQ, SQS.

## Features

- Kafka support
- RabbitMQ support
- Amazon SQS
- Redis queues
- Dead letter queues
- Message retry
- Consumer groups

## Usage

```typescript
import { Queue } from '@megaverse/queue';

const queue = new Queue({
  broker: 'kafka',
  brokers: ['localhost:9092'],
});

// Publish
await queue.publish('user-events', { type: 'created', userId: '123' });

// Consume
await queue.subscribe('user-events', async (message) => {
  await processEvent(message);
});
```
