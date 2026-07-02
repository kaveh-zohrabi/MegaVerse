# MegaVerse Makefile

.PHONY: help setup dev build test lint clean docker-up docker-down

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Install dependencies
	npm install

dev: ## Start frontend development
	cd apps/web && npm run dev

build: ## Build frontend
	cd apps/web && npm run build

test: ## Run tests
	npm test

lint: ## Lint code
	cd apps/web && npm run lint

clean: ## Clean build artifacts
	rm -rf node_modules apps/web/node_modules apps/web/.next

docker-up: ## Start MySQL
	docker-compose -f docker-compose.dev.yml up -d

docker-down: ## Stop MySQL
	docker-compose -f docker-compose.dev.yml down
