# MegaVerse Design Tokens

Design tokens for colors, typography, spacing, and more.

## Tokens

### Colors
- Primary: `#6366f1` (Indigo)
- Secondary: `#ec4899` (Pink)
- Success: `#22c55e` (Green)
- Warning: `#f59e0b` (Amber)
- Error: `#ef4444` (Red)

### Typography
- Font Family: Inter, system-ui
- Font Sizes: xs, sm, base, lg, xl, 2xl, 3xl
- Font Weights: normal, medium, semibold, bold

### Spacing
- 0: 0
- 1: 0.25rem
- 2: 0.5rem
- 3: 0.75rem
- 4: 1rem
- 6: 1.5rem
- 8: 2rem

## Usage

```css
@import '@megaverse/design-tokens/tokens/variables.css';

.button {
  background-color: var(--color-primary);
  padding: var(--spacing-4);
  font-size: var(--font-size-base);
}
```
