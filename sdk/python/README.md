# MegaVerse Python SDK

Official Python SDK for the MegaVerse API.

## Installation

```bash
pip install megaverse-sdk
```

## Quick Start

```python
from megaverse import MegaVerseClient

client = MegaVerseClient(
    api_key="your-api-key",
    base_url="https://api.megaverse.dev",
)

# Get user
user = client.users.get("user-id")

# Create post
post = client.posts.create(content="Hello MegaVerse!")
```

## Features

- Async support
- Type hints
- Automatic retries
- Pydantic models
- Context manager support

## Documentation

See [docs.megaverse.dev/sdk/python](https://docs.megaverse.dev/sdk/python)
