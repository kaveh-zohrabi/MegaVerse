# MegaVerse

A modular, polyglot software ecosystem built for scale.

## Quick Start

```bash
# Start MySQL
docker-compose -f docker-compose.dev.yml up -d

# Install dependencies
npm install

# Run frontend
cd apps/web && npm run dev
```

Open **http://localhost:3000**

### Prerequisites

- Node.js 18+
- npm
- Docker & Docker Compose

## Project Structure

```
megaverse/
├── apps/              # Frontend applications (Next.js)
├── services/          # Backend microservices (Go, Java, Python)
├── libraries/         # Shared libraries
├── shared/            # Shared types and constants
├── sdk/               # Client SDKs (JavaScript, Python, Go)
├── ai/                # AI/ML pipeline
├── gateway/           # API Gateway
├── infra/             # Infrastructure
├── docs/              # Documentation
├── tests/             # Test suites
└── scripts/           # Build and utility scripts
```

## Services

| Service | Language | Port | Description |
|---------|----------|------|-------------|
| api-gateway | Go | 8080 | Request routing, rate limiting |
| auth-service | Go | 8081 | Authentication, JWT |
| user-service | Go | 8082 | User profiles, follows |
| social-service | Java | 8083 | Posts, comments |
| messaging-service | Go | 8084 | Real-time messaging |
| ai-service | Python | 8085 | ML inference, embeddings |

## Tech Stack

| Layer | Technologies |
|-------|-------------|
| Frontend | Next.js, React, TypeScript, Tailwind |
| Backend | Go, Java, Python |
| Database | MySQL |
| AI/ML | Python, NumPy |

## Documentation

- [Architecture](./ARCHITECTURE.md)
- [Roadmap](./ROADMAP.md)
- [Contributing](./CONTRIBUTING.md)

## License

[MIT](./LICENSE)
