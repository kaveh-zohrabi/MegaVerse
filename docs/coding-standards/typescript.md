# TypeScript Coding Standards

## Style

- Use strict TypeScript (`strict: true`)
- Prefer interfaces over types for objects
- Use `const` by default, `let` when necessary
- Never use `var`
- Use template literals over string concatenation

## Naming Conventions

- `PascalCase` for:
  - Classes
  - Interfaces
  - Types
  - Enums
  - React components

- `camelCase` for:
  - Variables
  - Functions
  - Methods
  - Parameters
  - Properties

- `UPPER_SNAKE_CASE` for:
  - Constants
  - Environment variables
  - Enum values

- `kebab-case` for:
  - File names
  - CSS classes
  - URLs

## Functions

- Keep functions small and focused
- Use named functions over arrow functions for top-level
- Arrow functions for callbacks and inline functions
- Prefer async/await over .then()

## Error Handling

```typescript
try {
  const result = await fetchData();
  return result;
} catch (error) {
  if (error instanceof ValidationError) {
    throw new BadRequestError(error.message);
  }
  throw new InternalError('Failed to fetch data');
}
```

## Imports

```typescript
// External imports first
import express from 'express';
import { v4 as uuid } from 'uuid';

// Internal imports
import { UserService } from './user-service';
import { logger } from '@megaverse/logger';
```
