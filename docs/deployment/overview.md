# Deployment Documentation

## Overview

MegaVerse supports multiple deployment strategies.

## Docker Deployment

### Building Images

```bash
# Build all services
docker-compose build

# Build specific service
docker-compose build auth-service
```

### Running

```bash
# Development
docker-compose -f docker-compose.dev.yml up

# Production
docker-compose -f docker-compose.prod.yml up -d
```

## Kubernetes Deployment

### Prerequisites

- Kubernetes cluster
- kubectl configured
- Helm 3+

### Deploying

```bash
# Deploy to dev
helm install megaverse ./helm/megaverse -f ./helm/megaverse/values-dev.yaml

# Deploy to production
helm install megaverse ./helm/megaverse -f ./helm/megaverse/values-prod.yaml
```

### Scaling

```bash
# Manual scaling
kubectl scale deployment auth-service --replicas=3

# Auto-scaling
kubectl apply -f k8s/services/auth-service/hpa.yaml
```

## CI/CD Pipeline

### GitHub Actions

- **CI**: Lint, test, build on every push
- **CD**: Deploy on tag creation
- **Security**: Scan on schedule

### Deployment Flow

1. Code pushed to `develop`
2. CI runs tests
3. Merge to `main`
4. Tag release
5. CD deploys to staging
6. Manual approval
7. Deploy to production

## Rollback

```bash
# Helm rollback
helm rollback megaverse <revision>

# Kubernetes rollout undo
kubectl rollout undo deployment/auth-service
```

## Blue-Green Deployment

1. Deploy new version to green environment
2. Run smoke tests
3. Switch traffic to green
4. Keep blue for rollback

## Canary Deployment

1. Deploy to small percentage of traffic
2. Monitor metrics
3. Gradually increase traffic
4. Full deployment if healthy
