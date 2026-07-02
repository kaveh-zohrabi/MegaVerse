# MegaVerse Rust SDK

Official Rust SDK for the MegaVerse API.

## Installation

```toml
[dependencies]
megaverse-sdk = "0.1.0"
```

## Quick Start

```rust
use megaverse_sdk::MegaVerseClient;

#[tokio::main]
async fn main() {
    let client = MegaVerseClient::new("your-api-key")
        .with_base_url("https://api.megaverse.dev");

    let user = client.users().get("user-id").await.unwrap();
    println!("{}", user.name);
}
```

## Features

- Async/await support
- Type-safe API
- Automatic retries
- Connection pooling
- Zero-cost abstractions

## Documentation

See [docs.megaverse.dev/sdk/rust](https://docs.megaverse.dev/sdk/rust)
