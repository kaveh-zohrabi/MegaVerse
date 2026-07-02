# AI Service

Python service providing ML inference and embeddings.

## Tech Stack

- **Language**: Python 3.12
- **Framework**: FastAPI
- **ML**: NumPy (demonstration)
- **API**: OpenAPI/Swagger

## Features

- Text embedding generation
- Vector similarity search
- Document storage and retrieval
- In-memory vector store

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | / | Service info |
| GET | /health | Health check |
| POST | /embeddings | Generate embedding |
| POST | /search | Similarity search |
| POST | /store | Store document |
| DELETE | /store/{id} | Delete document |
| GET | /stats | Store statistics |

## API Documentation

Visit `http://localhost:8085/docs` for interactive API docs.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| PORT | 8085 | Server port |

## Development

```bash
# Install dependencies
pip install -e ".[dev]"

# Run server
uvicorn src.main:app --reload --port 8085

# Run tests
pytest
```

## Production Notes

The current embedding service uses hash-based vectors for demonstration.
Replace `src/services/embedding.py` with a real model:
- sentence-transformers
- OpenAI embeddings
- Custom trained model
