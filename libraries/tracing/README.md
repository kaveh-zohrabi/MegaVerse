# MegaVerse Tracing

Distributed tracing with OpenTelemetry.

## Features

- Span creation
- Context propagation
- Trace sampling
- Service graphs
- Performance analysis

## Usage

```typescript
import { Tracer } from '@megaverse/tracing';

const tracer = new Tracer({ service: 'user-service' });

const span = tracer.startSpan('get-user');
try {
  const user = await getUser(id);
  span.setStatus({ code: 'OK' });
  return user;
} catch (error) {
  span.setStatus({ code: 'ERROR', message: error.message });
  throw error;
} finally {
  span.end();
}
```
