# MegaVerse Validation Library

Request and data validation with schemas.

## Features

- Schema validation
- Request validation
- Type coercion
- Custom validators
- Error formatting
- i18n support

## Usage

```typescript
import { validate, z } from '@megaverse/validation';

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

const result = validate(schema, req.body);
if (!result.success) {
  return res.status(400).json(result.errors);
}
```
