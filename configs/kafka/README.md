# Kafka Configuration

## Overview

Apache Kafka configuration for event streaming.

## Topics

| Topic | Partitions | Replication | Retention |
|-------|-----------|-------------|-----------|
| user-events | 12 | 3 | 7 days |
| post-events | 12 | 3 | 7 days |
| order-events | 6 | 3 | 30 days |
| notification-events | 6 | 3 | 7 days |
| analytics-events | 24 | 3 | 90 days |

## Configuration

```properties
# Broker
broker.id=1
listeners=PLAINTEXT://0.0.0.0:9092
log.dirs=/var/lib/kafka/data

# Replication
default.replication.factor=3
min.insync.replicas=2

# Retention
log.retention.hours=168
log.retention.bytes=1073741824

# Performance
num.network.threads=8
num.io.threads=16
```

## Usage

```typescript
import { Kafka } from 'kafkajs';

const kafka = new Kafka({
  clientId: 'megaverse',
  brokers: ['localhost:9092'],
});

const producer = kafka.producer();
await producer.send({
  topic: 'user-events',
  messages: [{ value: JSON.stringify({ type: 'created', userId: '123' }) }],
});
```
