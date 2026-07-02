from pydantic import BaseModel
from typing import Optional, List
from datetime import datetime


class EmbeddingRequest(BaseModel):
    text: str
    model: str = "default"


class EmbeddingResponse(BaseModel):
    embedding: List[float]
    model: str
    dimensions: int


class SearchRequest(BaseModel):
    query: str
    limit: int = 10
    threshold: float = 0.7


class SearchResult(BaseModel):
    id: str
    content: str
    score: float
    metadata: Optional[dict] = None


class SearchResponse(BaseModel):
    results: List[SearchResult]
    query: str
    total: int


class HealthResponse(BaseModel):
    status: str
    service: str
    version: str
