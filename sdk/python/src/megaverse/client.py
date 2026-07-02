from typing import Optional
import httpx

from .models import User, Post, AuthResponse, PaginationParams


class MegaVerseClient:
    """Official Python SDK for MegaVerse API."""

    def __init__(self, base_url: str = "http://localhost:8080", api_key: Optional[str] = None):
        self.base_url = base_url.rstrip("/")
        self.token: Optional[str] = None
        self.headers = {"Content-Type": "application/json"}
        if api_key:
            self.headers["X-API-Key"] = api_key

    def _get_client(self) -> httpx.Client:
        headers = self.headers.copy()
        if self.token:
            headers["Authorization"] = f"Bearer {self.token}"
        return httpx.Client(base_url=self.base_url, headers=headers, timeout=30.0)

    def set_token(self, token: str) -> None:
        self.token = token

    # Auth
    def register(self, email: str, name: str, password: str) -> User:
        with self._get_client() as client:
            resp = client.post("/auth/register", json={"email": email, "name": name, "password": password})
            resp.raise_for_status()
            data = resp.json()
            return User(**data)

    def login(self, email: str, password: str) -> AuthResponse:
        with self._get_client() as client:
            resp = client.post("/auth/login", json={"email": email, "password": password})
            resp.raise_for_status()
            data = resp.json()
            self.set_token(data["access_token"])
            return AuthResponse(**data)

    # Users
    def get_user(self, user_id: str) -> User:
        with self._get_client() as client:
            resp = client.get(f"/users/{user_id}")
            resp.raise_for_status()
            return User(**resp.json())

    def update_profile(self, **kwargs) -> dict:
        with self._get_client() as client:
            resp = client.put("/users/me", json=kwargs)
            resp.raise_for_status()
            return resp.json()

    def follow(self, user_id: str) -> dict:
        with self._get_client() as client:
            resp = client.post(f"/users/{user_id}/follow")
            resp.raise_for_status()
            return resp.json()

    def unfollow(self, user_id: str) -> dict:
        with self._get_client() as client:
            resp = client.post(f"/users/{user_id}/unfollow")
            resp.raise_for_status()
            return resp.json()

    # Posts
    def create_post(self, content: str, media_url: Optional[str] = None) -> Post:
        with self._get_client() as client:
            payload = {"content": content}
            if media_url:
                payload["media_url"] = media_url
            resp = client.post("/posts", json=payload)
            resp.raise_for_status()
            return Post(**resp.json())

    def get_post(self, post_id: str) -> Post:
        with self._get_client() as client:
            resp = client.get(f"/posts/{post_id}")
            resp.raise_for_status()
            return Post(**resp.json())

    def delete_post(self, post_id: str) -> None:
        with self._get_client() as client:
            resp = client.delete(f"/posts/{post_id}")
            resp.raise_for_status()

    def get_feed(self, page: int = 1, limit: int = 20) -> list[Post]:
        with self._get_client() as client:
            resp = client.get("/feed", params={"page": page, "limit": limit})
            resp.raise_for_status()
            data = resp.json()
            return [Post(**p) for p in data.get("posts", [])]

    # Health
    def health(self) -> dict:
        with self._get_client() as client:
            resp = client.get("/health")
            resp.raise_for_status()
            return resp.json()
