# MegaVerse Makefile
# Root build orchestration

.PHONY: help setup dev build test lint clean deploy

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

setup: ## Install all dependencies
	@echo "Setting up MegaVerse..."
	bash scripts/setup/install-deps.sh
	bash scripts/setup/setup-env.sh
	bash scripts/setup/setup-db.sh

dev: ## Start development environment
	docker-compose -f docker-compose.dev.yml up -d
	bash scripts/development/dev.sh

build: ## Build all packages
	bash build/scripts/build.sh

test: ## Run all tests
	bash scripts/development/test.sh

lint: ## Run all linters
	bash scripts/development/lint.sh

format: ## Format all code
	bash scripts/development/format.sh

clean: ## Clean build artifacts
	bash build/scripts/cleanup.sh

deploy: ## Deploy to production
	bash scripts/deployment/deploy.sh

migrate: ## Run database migrations
	bash scripts/database/migrate.sh

seed: ## Seed database
	bash scripts/database/seed.sh

backup: ## Backup database
	bash scripts/database/backup.sh

logs: ## View service logs
	bash scripts/deployment/logs.sh

status: ## Check service status
	bash scripts/deployment/status.sh

healthcheck: ## Run health checks
	bash build/scripts/healthcheck.sh

docker-build: ## Build Docker images
	bash build/packages/docker/build.sh

k8s-deploy: ## Deploy to Kubernetes
	bash scripts/deployment/deploy.sh k8s

terraform-init: ## Initialize Terraform
	cd infrastructure/terraform/environments/dev && terraform init

terraform-plan: ## Plan Terraform changes
	cd infrastructure/terraform/environments/dev && terraform plan

terraform-apply: ## Apply Terraform changes
	cd infrastructure/terraform/environments/dev && terraform apply

ai-train: ## Train AI models
	bash ai/scripts/train.py

ai-evaluate: ## Evaluate AI models
	bash ai/scripts/evaluate.py

ai-deploy: ## Deploy AI models
	bash ai/scripts/deploy.py
