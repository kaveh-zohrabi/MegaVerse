# MegaVerse Metrics

Application metrics collection and export.

## Features

- Counter, Gauge, Histogram
- Prometheus export
- Custom metrics
- Labels/tags
- Metric aggregation

## Usage

```typescript
import { Metrics } from '@megaverse/metrics';

const metrics = new Metrics({ prefix: 'megaverse' });

// Counter
metrics.counter('requests_total').inc({ method: 'GET', path: '/api' });

// Histogram
metrics.histogram('request_duration').observe(duration);
```
