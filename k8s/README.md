# MegaVerse Kubernetes Manifests

## Overview

Kubernetes deployment manifests for all services.

## Structure

```
k8s/
├── base/               # Base configurations
├── services/           # Service deployments
├── databases/          # Database deployments
├── infrastructure/     # Infrastructure components
├── environments/       # Environment-specific configs
└── overlays/           # Kustomize overlays
```

## Usage

### Base Deployment

```bash
kubectl apply -f k8s/base/
```

### Service Deployment

```bash
kubectl apply -f k8s/services/auth-service/
```

### Environment Overlay

```bash
kubectl apply -k k8s/overlays/dev/
kubectl apply -k k8s/overlays/staging/
kubectl apply -k k8s/overlays/prod/
```

## Components

### Services
- Deployments
- Services
- Ingress
- ConfigMaps
- Secrets
- HPA (Horizontal Pod Autoscaler)
- PDB (Pod Disruption Budget)

### Databases
- StatefulSets
- PersistentVolumeClaims
- Services
- Headless Services

### Infrastructure
- Ingress controllers
- Certificates
- Monitoring
- Logging
- DNS

## Best Practices

- Use namespaces for isolation
- Resource limits and requests
- Health checks (liveness, readiness)
- Security contexts
- Network policies
