# Go Coding Standards

## Style

- Follow Effective Go guidelines
- Use `gofmt` for formatting
- Use `golangci-lint` for linting

## Naming Conventions

- `CamelCase` for exported names
- `camelCase` for unexported names
- `UPPER_SNAKE_CASE` for constants
- Short names for short scopes

## Error Handling

```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}
```

## Project Layout

```
cmd/           # Main applications
internal/      # Private application code
pkg/           # Public library code
api/           # API definitions
configs/       # Configuration files
```

## Concurrency

- Use goroutines for concurrent operations
- Use channels for communication
- Use sync.WaitGroup for synchronization
- Use context for cancellation

## Testing

```go
func TestSomething(t *testing.T) {
    // Arrange
    input := "test"
    
    // Act
    result := doSomething(input)
    
    // Assert
    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}
```
