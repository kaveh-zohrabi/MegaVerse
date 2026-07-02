# MegaVerse Events Library

Event-driven messaging with Kafka and RabbitMQ.

## Features

- Event publishing
- Event consumption
- Event schema validation
- Dead letter queues
- Event replay
- Event sourcing support

## Usage

```typescript
import { EventBus } from '@megaverse/events';

const bus = new EventBus({ broker: 'kafka' });

// Publish
await bus.publish('user.created', { userId: '123', email: 'user@example.com' });

// Subscribe
bus.subscribe('user.created', async (event) => {
  await sendWelcomeEmail(event.data.email);
});
```
