# MegaVerse Java SDK

Official Java SDK for the MegaVerse API.

## Installation

```xml
<dependency>
    <groupId>dev.megaverse</groupId>
    <artifactId>megaverse-java-sdk</artifactId>
    <version>0.1.0</version>
</dependency>
```

## Quick Start

```java
import dev.megaverse.sdk.MegaVerseClient;

MegaVerseClient client = MegaVerseClient.builder()
    .apiKey("your-api-key")
    .baseUrl("https://api.megaverse.dev")
    .build();

User user = client.users().get("user-id");
```

## Features

- Builder pattern
- Automatic retries
- Type-safe API
- Async support
- Connection pooling

## Documentation

See [docs.megaverse.dev/sdk/java](https://docs.megaverse.dev/sdk/java)
