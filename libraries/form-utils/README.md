# MegaVerse Form Utils

Form handling and validation utilities.

## Features

- Form validation
- Field manipulation
- Error formatting
- Form state management
- Dirty/pristine tracking

## Usage

```typescript
import { useForm, validate } from '@megaverse/form-utils';

const form = useForm({
  initialValues: { email: '', password: '' },
  validate: (values) => ({
    email: !values.email ? 'Required' : undefined,
    password: values.password.length < 8 ? 'Too short' : undefined,
  }),
  onSubmit: async (values) => {
    await login(values);
  },
});
```
