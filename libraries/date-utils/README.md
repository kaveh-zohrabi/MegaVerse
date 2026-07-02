# MegaVerse Date Utils

Date and time utilities.

## Features

- Date formatting
- Relative time
- Timezone handling
- Duration calculations
- Business hours
- Holiday support

## Usage

```typescript
import { formatDate, relativeTime, addDays } from '@megaverse/date-utils';

formatDate(new Date(), 'yyyy-MM-dd'); // "2026-07-01"
relativeTime(new Date(Date.now() - 3600000)); // "1 hour ago"
addDays(new Date(), 7); // Date 7 days from now
```
