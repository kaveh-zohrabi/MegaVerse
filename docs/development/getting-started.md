# Development Guide

## Getting Started

### Prerequisites

- Node.js 20+
- Go 1.22+
- Python 3.12+
- Docker & Docker Compose
- pnpm 9+

### Setup

```bash
# Clone repository
git clone https://github.com/your-org/megaverse.git
cd megaverse

# Install dependencies
make setup

# Start development environment
make dev
```

### Project Structure

```
megaverse/
├── apps/           # Frontend applications
├── services/       # Backend microservices
├── packages/       # Shared configurations
├── libraries/      # Reusable libraries
├── shared/         # Shared models and types
├── sdk/            # Client SDKs
├── ai/             # AI/ML pipeline
├── cloud/          # Cloud infrastructure
├── gateway/        # API Gateway
├── docs/           # Documentation
├── tests/          # Test suites
├── configs/        # Configuration files
├── build/          # Build system
├── scripts/        # Utility scripts
└── tools/          # Developer tools
```

### Running Services

```bash
# Start all services
docker-compose up -d

# Start specific service
docker-compose up auth-service

# View logs
docker-compose logs -f auth-service
```

### Database

```bash
# Run migrations
make migrate

# Seed database
make seed

# Backup database
make backup
```

### Testing

```bash
# Run all tests
make test

# Run specific test suite
pnpm test
go test ./...
pytest
```

### Code Quality

```bash
# Lint
make lint

# Format
make format

# Type check
pnpm typecheck
```
