# Coding Standards

## General Principles

1. **Clarity over cleverness**
2. **DRY (Don't Repeat Yourself)**
3. **YAGNI (You Aren't Gonna Need It)**
4. **Boy Scout Rule**: Leave code better than you found it
5. **Single Responsibility**: One function, one job

## Code Review Checklist

- [ ] Code follows style guidelines
- [ ] Tests are included
- [ ] Documentation is updated
- [ ] No security vulnerabilities
- [ ] Performance considerations
- [ ] Error handling is appropriate
- [ ] Logging is adequate

## Git Conventions

### Commit Messages

```
type(scope): description

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `style`: Formatting
- `refactor`: Code refactoring
- `test`: Tests
- `chore`: Maintenance

### Branch Naming

```
feature/TICKET-123-user-authentication
bugfix/TICKET-456-fix-login-error
hotfix/TICKET-789-security-patch
```

## Language-Specific Standards

See:
- [TypeScript](./typescript.md)
- [Go](./go.md)
- [Python](./python.md)
- [Rust](./rust.md)
- [Java](./java.md)
- [Kotlin](./kotlin.md)
- [Dart](./dart.md)
