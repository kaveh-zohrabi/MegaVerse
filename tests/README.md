# MegaVerse Tests

## Overview

This directory contains all test suites for MegaVerse.

## Structure

```
tests/
├── unit/           # Unit tests for individual functions
├── integration/    # Integration tests for service interactions
└── e2e/            # End-to-end tests for full workflows
```

## Running Tests

```bash
# All tests
pnpm test

# Go service tests
make test-go

# Python service tests
make test-python

# TypeScript tests
make test-typescript
```

## Writing Tests

- Place unit tests next to the code they test
- Use descriptive test names
- Follow AAA pattern: Arrange, Act, Assert
- Mock external dependencies
