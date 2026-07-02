# Java Coding Standards

## Style

- Follow Google Java Style Guide
- Use Spotless for formatting
- Use Checkstyle for style checking

## Naming Conventions

- `PascalCase` for:
  - Classes
  - Interfaces
  - Enums

- `camelCase` for:
  - Methods
  - Variables
  - Parameters

- `UPPER_SNAKE_CASE` for:
  - Constants

## Project Structure

```
src/main/java/
├── dev/megaverse/
│   ├── controllers/
│   ├── services/
│   ├── repositories/
│   ├── models/
│   ├── dto/
│   ├── config/
│   └── utils/
```

## Error Handling

```java
try {
    User user = userService.getUser(id);
    return ResponseEntity.ok(user);
} catch (UserNotFoundException e) {
    return ResponseEntity.notFound().build();
} catch (Exception e) {
    return ResponseEntity.internalServerError().build();
}
```

## Testing

```java
@SpringBootTest
class UserServiceTest {
    @Autowired
    private UserService userService;

    @Test
    void shouldGetUser() {
        // Arrange
        String userId = "123";

        // Act
        User user = userService.getUser(userId);

        // Assert
        assertNotNull(user);
        assertEquals(userId, user.getId());
    }
}
```
