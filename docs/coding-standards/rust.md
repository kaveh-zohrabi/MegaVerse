# Rust Coding Standards

## Style

- Follow Rust API Guidelines
- Use `clippy` for linting
- Use `rustfmt` for formatting

## Naming Conventions

- `snake_case` for:
  - Variables
  - Functions
  - Modules
  - Macros

- `PascalCase` for:
  - Types
  - Traits
  - Enums

- `SCREAMING_SNAKE_CASE` for:
  - Constants
  - Statics

## Error Handling

```rust
fn process_data(data: &Data) -> Result<Output, Error> {
    let result = parse_input(data)?;
    let output = transform(result)?;
    Ok(output)
}
```

## Ownership

- Prefer borrowing over cloning
- Use `&str` over `&String`
- Use `Cow<str>` for flexible string handling
- Use `Arc` for shared ownership across threads

## Testing

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process_data() {
        let data = Data::new("test");
        let result = process_data(&data);
        assert!(result.is_ok());
    }
}
```

## Documentation

```rust
/// Processes the input data and returns the result.
///
/// # Arguments
///
/// * `data` - The input data to process
///
/// # Returns
///
/// The processed output or an error.
///
/// # Examples
///
/// ```
/// let result = process_data(&data).unwrap();
/// ```
pub fn process_data(data: &Data) -> Result<Output, Error> {
    // ...
}
```
