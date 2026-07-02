# MegaVerse Go SDK

Official Go SDK for the MegaVerse API.

## Installation

```bash
go get megaverse.dev/sdk/go
```

## Quick Start

```go
package main

import (
    "fmt"
    megaverse "megaverse.dev/sdk/go"
)

func main() {
    client := megaverse.NewClient("http://localhost:8080", "")

    // Login
    auth, _ := client.Login("user@example.com", "password")
    fmt.Println("Token:", auth.AccessToken)

    // Get user
    user, _ := client.GetUser("user-id")
    fmt.Println("User:", user.Name)

    // Create post
    post, _ := client.CreatePost("Hello MegaVerse!")
    fmt.Println("Post:", post.ID)
}
```

## API Reference

### Client

```go
NewClient(baseURL, apiKey string) *Client
```

### Auth

- `Register(email, name, password string) (*User, error)`
- `Login(email, password string) (*AuthResponse, error)`
- `SetToken(token string)`

### Users

- `GetUser(id string) (*User, error)`
- `Follow(userID string) error`
- `Unfollow(userID string) error`

### Posts

- `CreatePost(content string) (*Post, error)`
- `GetPost(id string) (*Post, error)`
- `DeletePost(id string) error`
- `GetFeed(params *PaginationParams) ([]Post, error)`

### Health

- `Health() (map[string]string, error)`
