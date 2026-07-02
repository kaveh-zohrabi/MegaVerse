from fastapi import APIRouter, HTTPException
from ..models.schemas import (
    EmbeddingRequest,
    EmbeddingResponse,
    SearchRequest,
    SearchResponse,
    HealthResponse,
)
from ..services.embedding import EmbeddingService
from ..services.vector_store import VectorStore

router = APIRouter()

# Initialize services
embedding_service = EmbeddingService(dimensions=128)
vector_store = VectorStore(embedding_service)


@router.get("/health", response_model=HealthResponse)
async def health_check():
    return HealthResponse(
        status="healthy",
        service="ai-service",
        version="0.1.0",
    )


@router.post("/embeddings", response_model=EmbeddingResponse)
async def create_embedding(request: EmbeddingRequest):
    """Generate embedding for text."""
    try:
        embedding = embedding_service.generate_embedding(request.text)
        return EmbeddingResponse(
            embedding=embedding,
            model=request.model,
            dimensions=len(embedding),
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.post("/search", response_model=SearchResponse)
async def search(request: SearchRequest):
    """Search for similar content."""
    try:
        results = vector_store.search(
            query=request.query,
            limit=request.limit,
            threshold=request.threshold,
        )
        return SearchResponse(
            results=results,
            query=request.query,
            total=len(results),
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))


@router.post("/store")
async def store_document(id: str, content: str, metadata: dict = None):
    """Add document to vector store."""
    vector_store.add(id, content, metadata)
    return {"status": "stored", "id": id}


@router.delete("/store/{id}")
async def delete_document(id: str):
    """Delete document from vector store."""
    if vector_store.delete(id):
        return {"status": "deleted", "id": id}
    raise HTTPException(status_code=404, detail="Document not found")


@router.get("/stats")
async def get_stats():
    """Get vector store statistics."""
    return {
        "documents": vector_store.count(),
        "dimensions": embedding_service.dimensions,
    }
