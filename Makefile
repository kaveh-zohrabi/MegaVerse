# MegaVerse Makefile

.PHONY: help setup dev build test lint format clean docker-up docker-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Install dependencies
	pnpm install

dev: ## Start development
	pnpm dev

build: ## Build all packages
	pnpm build

test: ## Run tests
	pnpm test

lint: ## Lint code
	pnpm lint

format: ## Format code
	pnpm format

clean: ## Clean build artifacts
	pnpm clean

docker-up: ## Start MySQL and Redis
	docker-compose -f docker-compose.dev.yml up -d

docker-down: ## Stop MySQL and Redis
	docker-compose -f docker-compose.dev.yml down

docker-build: ## Build Docker images
	docker-compose build

db-migrate: ## Run database migrations (service-specific)
	@echo "Run: cd services/<service> && make migrate"

db-seed: ## Seed database (service-specific)
	@echo "Run: cd services/<service> && make seed"

test-go: ## Run Go service tests
	@for dir in services/*/; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "Testing $$dir..."; \
			cd $$dir && go test ./... && cd ../..; \
		fi; \
	done

test-python: ## Run Python service tests
	cd services/ai-service && python -m pytest

test-typescript: ## Run TypeScript tests
	pnpm test
