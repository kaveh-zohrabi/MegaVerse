# Python Coding Standards

## Style

- Follow PEP 8
- Use `ruff` for linting
- Use `black` for formatting
- Type hints for all functions

## Naming Conventions

- `snake_case` for:
  - Variables
  - Functions
  - Methods
  - Modules

- `PascalCase` for:
  - Classes
  - Exceptions

- `UPPER_SNAKE_CASE` for:
  - Constants

## Type Hints

```python
def process_user(user: User) -> dict[str, Any]:
    """Process user data."""
    return {"id": user.id, "name": user.name}
```

## Error Handling

```python
try:
    result = await process_data(data)
except ValidationError as e:
    logger.warning(f"Validation failed: {e}")
    raise BadRequestError(str(e))
except Exception as e:
    logger.error(f"Unexpected error: {e}")
    raise InternalError("Processing failed")
```

## Async/Await

```python
async def get_user(user_id: str) -> User:
    async with aiohttp.ClientSession() as session:
        async with session.get(f"/users/{user_id}") as response:
            return await response.json()
```

## Testing

```python
@pytest.mark.asyncio
async def test_get_user():
    # Arrange
    user_id = "123"
    
    # Act
    user = await get_user(user_id)
    
    # Assert
    assert user.id == user_id
```
