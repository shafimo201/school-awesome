# School ERP SaaS Platform

Enterprise-grade School ERP targeting 1000+ schools. This repository contains a Go backend designed for production deployment on AWS EKS with PostgreSQL, Redis, and observability.

## Architecture

- Clean Architecture (ports & adapters)
- Multi-tenancy by `school_id`
- Layered separation: domain, use cases, adapters, frameworks
- Independent modules for auth, school, user, audit, notifications

## Folder Structure

- `cmd/` — application entrypoints
- `internal/` — private implementation details
- `pkg/` — reusable packages and shared contracts
- `api/` — OpenAPI specs and API definitions
- `configs/` — environment and deployment configs
- `deployments/` — Helm charts and K8s manifests
- `scripts/` — utility scripts for dev and infra
- `migrations/` — Goose DB migrations

## Getting Started

1. Install Go 1.24+
2. Install dependencies: `go mod tidy`
3. Run local server: `go run ./cmd/server`

## Production Requirements

- JWT authentication with refresh tokens
- RBAC with permission enforcement
- Audit logging and soft-deletes
- Observability: Prometheus + OpenTelemetry
- AWS support: EKS, RDS, ECR, S3, Secrets Manager

## Development Rules

- No unnecessary code generation
- Every module is independently testable
- Follow Go best practices and strict linting
