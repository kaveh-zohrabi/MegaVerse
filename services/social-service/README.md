# Social Service

Java service handling social features - posts, comments, and reactions.

## Tech Stack

- **Language**: Java 21
- **Framework**: Spring Boot 3.2
- **Database**: MySQL 8.0
- **ORM**: Spring Data JPA

## Features

- Create, read, delete posts
- Add comments to posts
- User feed generation
- Post ownership validation

## Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| GET | /health | No | Health check |
| POST | /posts | Yes | Create post |
| GET | /posts/{id} | No | Get post |
| DELETE | /posts/{id} | Yes | Delete post |
| GET | /users/{userId}/posts | No | Get user posts |
| GET | /feed | No | Get global feed |
| POST | /posts/{postId}/comments | Yes | Add comment |
| GET | /posts/{postId}/comments | No | Get comments |

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| server.port | 8083 | Server port |
| spring.datasource.url | jdbc:mysql://localhost:3306/megaverse | Database URL |

## Development

```bash
# Run
./mvnw spring-boot:run

# Build
./mvnw clean package
```

## Database

Run migration before starting:

```bash
mysql -u root -p megaverse < src/main/resources/migrations/001_create_posts.sql
```
