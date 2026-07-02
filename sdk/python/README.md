# MegaVerse Python SDK

Official Python SDK for the MegaVerse API.

## Installation

```bash
pip install megaverse-sdk
```

## Quick Start

```python
from megaverse import MegaVerseClient

client = MegaVerseClient(base_url="http://localhost:8080")

# Login
client.login("user@example.com", "password")

# Get user
user = client.get_user("user-id")

# Create post
post = client.create_post("Hello MegaVerse!")

# Get feed
posts = client.get_feed(limit=10)
```

## API Reference

### Client

```python
MegaVerseClient(base_url: str, api_key: str = None)
```

### Auth

- `register(email, name, password)` - Register new user
- `login(email, password)` - Login and get tokens
- `set_token(token)` - Set auth token

### Users

- `get_user(user_id)` - Get user profile
- `update_profile(**kwargs)` - Update own profile
- `follow(user_id)` - Follow user
- `unfollow(user_id)` - Unfollow user

### Posts

- `create_post(content, media_url=None)` - Create post
- `get_post(post_id)` - Get post
- `delete_post(post_id)` - Delete post
- `get_feed(page=1, limit=20)` - Get feed

### Health

- `health()` - Check API health
