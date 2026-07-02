# MegaVerse Localization

## Overview

Multi-language support for the platform.

## Supported Languages

| Code | Language | Status |
|------|----------|--------|
| en | English | Primary |
| es | Spanish | Full |
| fr | French | Full |
| de | German | Full |
| ja | Japanese | Full |
| ko | Korean | Full |
| zh | Chinese (Simplified) | Full |
| ar | Arabic | Partial |
| pt | Portuguese | Full |
| ru | Russian | Partial |

## Translation Files

```
configs/localization/
├── en/
│   ├── common.json
│   ├── auth.json
│   ├── dashboard.json
│   └── errors.json
├── es/
│   └── ...
└── ...
```

## Usage

```typescript
import { t } from '@megaverse/i18n';

const greeting = t('common.greeting', { name: 'John' });
// English: "Hello, John!"
// Spanish: "¡Hola, John!"
```

## Contributing Translations

1. Fork the repository
2. Add translations in `configs/localization/`
3. Submit PR
4. Review by native speakers
