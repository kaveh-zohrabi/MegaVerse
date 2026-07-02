# MegaVerse Swift SDK

Official Swift SDK for the MegaVerse API.

## Installation

### Swift Package Manager

```swift
dependencies: [
    .package(url: "https://github.com/megaverse/sdk-swift.git", from: "0.1.0")
]
```

## Quick Start

```swift
import MegaVerseSDK

let client = MegaVerseClient(apiKey: "your-api-key")

let user = try await client.users.get(id: "user-id")
print(user.name)
```

## Features

- Async/await support
- Combine publishers
- Type-safe API
- Automatic retries
- iOS/macOS support

## Documentation

See [docs.megaverse.dev/sdk/swift](https://docs.megaverse.dev/sdk/swift)
