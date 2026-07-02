import numpy as np
from typing import List
import hashlib
import json


class EmbeddingService:
    """Simple embedding service using numpy for demonstration.
    
    In production, this would use a proper embedding model like:
    - sentence-transformers
    - OpenAI embeddings
    - Custom trained model
    """

    def __init__(self, dimensions: int = 128):
        self.dimensions = dimensions

    def generate_embedding(self, text: str) -> List[float]:
        """Generate a simple embedding from text.
        
        This is a hash-based embedding for demonstration.
        Replace with actual ML model in production.
        """
        # Create deterministic hash-based embedding
        hash_obj = hashlib.sha256(text.encode())
        hash_bytes = hash_obj.digest()

        # Convert to float vector
        embedding = []
        for i in range(0, min(len(hash_bytes), self.dimensions)):
            embedding.append(hash_bytes[i] / 255.0)

        # Pad if needed
        while len(embedding) < self.dimensions:
            embedding.append(0.0)

        # Normalize
        norm = np.linalg.norm(embedding)
        if norm > 0:
            embedding = (np.array(embedding) / norm).tolist()

        return embedding[:self.dimensions]

    def cosine_similarity(self, a: List[float], b: List[float]) -> float:
        """Calculate cosine similarity between two vectors."""
        a_np = np.array(a)
        b_np = np.array(b)
        return float(np.dot(a_np, b_np) / (np.linalg.norm(a_np) * np.linalg.norm(b_np)))
