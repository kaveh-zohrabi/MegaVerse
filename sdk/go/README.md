# MegaVerse Go SDK

Official Go SDK for the MegaVerse API.

## Installation

```bash
go get megaverse.dev/sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    "megaverse.dev/sdk"
)

func main() {
    client := sdk.NewClient(
        sdk.WithAPIKey("your-api-key"),
        sdk.WithBaseURL("https://api.megaverse.dev"),
    )

    user, err := client.Users.Get("user-id")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(user.Name)
}
```

## Features

- Context support
- Automatic retries
- Type-safe API
- Streaming support
- Connection pooling

## Documentation

See [docs.megaverse.dev/sdk/go](https://docs.megaverse.dev/sdk/go)
