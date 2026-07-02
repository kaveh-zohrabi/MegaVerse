# MegaVerse Makefile

.PHONY: help setup dev build test lint format clean

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

docker-up: ## Start Docker services
	docker-compose -f docker-compose.dev.yml up -d

docker-down: ## Stop Docker services
	docker-compose -f docker-compose.dev.yml down

docker-build: ## Build Docker images
	docker-compose build

db-migrate: ## Run database migrations
	@echo "Run migration commands for your service"

db-seed: ## Seed database
	@echo "Run seed commands for your service"
