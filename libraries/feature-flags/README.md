# MegaVerse Feature Flags

Feature flag management for gradual rollouts.

## Features

- Boolean flags
- Percentage rollouts
- User targeting
- A/B testing
- Remote configuration
- Flag analytics

## Usage

```typescript
import { FeatureFlags } from '@megaverse/feature-flags';

const flags = new FeatureFlags({
  service: 'user-service',
  endpoint: 'https://flags.megaverse.dev',
});

if (await flags.isEnabled('new-signup-flow', { userId })) {
  // New flow
} else {
  // Old flow
}
```
