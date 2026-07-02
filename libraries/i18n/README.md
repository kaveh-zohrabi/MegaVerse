# MegaVerse i18n Library

Internationalization and localization support.

## Features

- Multiple language support
- Pluralization
- Number/date formatting
- Currency formatting
- RTL support
- Namespace loading

## Usage

```typescript
import { t, setLocale } from '@megaverse/i18n';

setLocale('en');
const greeting = t('welcome', { name: 'John' }); // "Welcome, John!"

setLocale('es');
const greetingEs = t('welcome', { name: 'John' }); // "¡Bienvenido, John!"
```

## Supported Languages

- English (en)
- Spanish (es)
- French (fr)
- German (de)
- Japanese (ja)
- Korean (ko)
- Chinese (zh)
- Arabic (ar)
- Portuguese (pt)
- Russian (ru)
