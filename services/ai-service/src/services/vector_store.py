from typing import List, Optional, Dict
from ..models.schemas import SearchResult
from .embedding import EmbeddingService


class VectorStore:
    """Simple in-memory vector store for demonstration.
    
    In production, use:
    - Pinecone
    - Weaviate
    - Milvus
    - Qdrant
    """

    def __init__(self, embedding_service: EmbeddingService):
        self.embedding_service = embedding_service
        self.documents: Dict[str, dict] = {}

    def add(self, id: str, content: str, metadata: Optional[dict] = None):
        """Add a document to the store."""
        embedding = self.embedding_service.generate_embedding(content)
        self.documents[id] = {
            "content": content,
            "embedding": embedding,
            "metadata": metadata or {},
        }

    def search(self, query: str, limit: int = 10, threshold: float = 0.7) -> List[SearchResult]:
        """Search for similar documents."""
        query_embedding = self.embedding_service.generate_embedding(query)

        results = []
        for doc_id, doc in self.documents.items():
            score = self.embedding_service.cosine_similarity(query_embedding, doc["embedding"])
            if score >= threshold:
                results.append(SearchResult(
                    id=doc_id,
                    content=doc["content"],
                    score=score,
                    metadata=doc["metadata"],
                ))

        # Sort by score
        results.sort(key=lambda x: x.score, reverse=True)
        return results[:limit]

    def delete(self, id: str) -> bool:
        """Delete a document from the store."""
        if id in self.documents:
            del self.documents[id]
            return True
        return False

    def count(self) -> int:
        """Get the number of documents."""
        return len(self.documents)
