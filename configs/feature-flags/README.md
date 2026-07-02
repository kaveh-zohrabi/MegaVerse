# MegaVerse Feature Flags

## Overview

Feature flag management for gradual rollouts and A/B testing.

## Usage

### Configuration

```json
{
  "flags": {
    "new-signup-flow": {
      "enabled": true,
      "percentage": 50,
      "variants": ["control", "variant-a", "variant-b"],
      "targeting": {
        "users": ["user-id-1", "user-id-2"],
        "countries": ["US", "CA"],
        "plans": ["pro", "enterprise"]
      }
    }
  }
}
```

### Code Usage

```typescript
import { FeatureFlags } from '@megaverse/feature-flags';

const flags = new FeatureFlags({
  service: 'user-service',
  endpoint: 'https://flags.megaverse.dev',
});

// Check flag
if (await flags.isEnabled('new-signup-flow', { userId })) {
  // New flow
}

// Get variant
const variant = await flags.getVariant('new-signup-flow', { userId });
```

## Environments

- Development: `feature-flags/development.json`
- Staging: `feature-flags/staging.json`
- Production: `feature-flags/production.json`
