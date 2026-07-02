# MegaVerse C# SDK

Official C# SDK for the MegaVerse API.

## Installation

```bash
dotnet add package MegaVerse.SDK
```

## Quick Start

```csharp
using MegaVerse.SDK;

var client = new MegaVerseClient("your-api-key");

var user = await client.Users.GetAsync("user-id");
Console.WriteLine(user.Name);
```

## Features

- Async/await support
- Type-safe API
- Automatic retries
- LINQ support
- Dependency injection ready

## Documentation

See [docs.megaverse.dev/sdk/csharp](https://docs.megaverse.dev/sdk/csharp)
