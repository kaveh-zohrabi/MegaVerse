# MegaVerse Deployment Guide

## Prerequisites

- Docker & Docker Compose
- Kubernetes cluster (production)
- Terraform (infrastructure)
- pnpm 9+
- Node.js 20+

## Local Development

```bash
# Install dependencies
pnpm install

# Start infrastructure
docker-compose up -d

# Start services
pnpm dev
```

## Docker Deployment

```bash
# Build images
docker-compose build

# Start all services
docker-compose up -d

# View logs
docker-compose logs -f
```

## Kubernetes Deployment

```bash
# Apply manifests
kubectl apply -f infra/terraform/

# Check status
kubectl get pods
kubectl get services
```

## Terraform

```bash
# Initialize
cd infra/terraform/environments/dev
terraform init

# Plan changes
terraform plan

# Apply changes
terraform apply
```

## Environment Variables

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
# Edit .env with your values
```

## Monitoring

```bash
# Prometheus: http://localhost:9090
# Grafana: http://localhost:3001
# Default Grafana login: admin/admin
```

## Database

```bash
# Run migrations
# (service-specific commands)

# Backup
pg_dump -U megaverse megaverse > backup.sql

# Restore
psql -U megaverse megaverse < backup.sql
```
