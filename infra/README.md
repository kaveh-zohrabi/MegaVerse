# MegaVerse Infrastructure

## Overview

Infrastructure as Code and deployment configurations.

## Structure

```
infra/
├── terraform/      # Terraform modules and environments
├── monitoring/     # Prometheus, Grafana configs
└── docker/         # Docker build contexts
```

## Environments

- **dev**: Local development
- **staging**: Pre-production testing
- **prod**: Production deployment

## Usage

```bash
# Terraform
cd infra/terraform/environments/dev
terraform init
terraform plan
terraform apply

# Docker
docker-compose -f infra/docker/docker-compose.yml up
```
