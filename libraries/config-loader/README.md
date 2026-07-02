# MegaVerse Config Loader

Configuration management with environment support.

## Features

- Environment variables
- Config files (JSON, YAML, TOML)
- Secret management
- Config validation
- Hot reloading
- Config merging

## Usage

```typescript
import { Config } from '@megaverse/config-loader';

const config = new Config({
  sources: ['env', './config.json'],
  schema: configSchema,
});

const dbUrl = config.get('database.url');
```
