# Contributing to MegaVerse

Thank you for your interest in contributing!

## Getting Started

1. Fork the repository
2. Clone your fork
3. Install dependencies: `pnpm install`
4. Create a feature branch: `git checkout -b feat/my-feature`
5. Make your changes
6. Submit a pull request

## Development

```bash
# Install dependencies
pnpm install

# Start all services
pnpm dev

# Run tests
pnpm test

# Lint code
pnpm lint

# Format code
pnpm format
```

## Project Layout

| Directory | Description |
|-----------|-------------|
| `apps/` | Frontend applications |
| `services/` | Backend microservices |
| `libraries/` | Shared libraries |
| `sdk/` | Client SDKs |
| `ai/` | AI/ML pipeline |
| `docs/` | Documentation |

## Code Standards

- TypeScript: strict mode, ESLint + Prettier
- Go: `gofmt`, `golangci-lint`
- Python: `ruff`, type hints required
- Java: Google Java Style, Spotless

## Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add user authentication
fix: resolve race condition in cache
docs: update API reference
test: add integration tests for payments
chore: update dependencies
```

## Pull Request Checklist

- [ ] Code follows project style
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] CI passes
- [ ] PR description explains changes

## Code Review

- Be respectful and constructive
- Focus on code quality and correctness
- Suggest improvements when possible

## Questions?

Open a [GitHub Discussion](https://github.com/megaverse/megaverse/discussions).
