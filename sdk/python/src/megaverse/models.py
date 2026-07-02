from dataclasses import dataclass
from typing import Optional, List


@dataclass
class User:
    id: str
    email: str
    name: str
    role: str


@dataclass
class Post:
    id: str
    user_id: str
    content: str
    media_url: Optional[str] = None
    likes_count: int = 0
    comments_count: int = 0
    created_at: str = ""


@dataclass
class AuthResponse:
    access_token: str
    refresh_token: str
    expires_at: int


@dataclass
class PaginationParams:
    page: int = 1
    limit: int = 20
