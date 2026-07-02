# MegaVerse Serialization

JSON/Protobuf/MessagePack serialization.

## Features

- JSON serialization
- Protocol Buffers
- MessagePack
- Custom serializers
- Schema validation
- Versioning

## Usage

```typescript
import { Serializer } from '@megaverse/serialization';

const serializer = new Serializer({ format: 'json' });

// Serialize
const data = serializer.serialize({ key: 'value' });

// Deserialize
const obj = serializer.deserialize(data);
```
