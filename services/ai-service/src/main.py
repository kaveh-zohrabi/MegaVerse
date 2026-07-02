from fastapi import FastAPI
from .api.routes import router

app = FastAPI(
    title="MegaVerse AI Service",
    description="AI/ML inference and embeddings service",
    version="0.1.0",
)

app.include_router(router)


@app.get("/")
async def root():
    return {
        "service": "ai-service",
        "version": "0.1.0",
        "docs": "/docs",
    }
