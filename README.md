# School Awesome

Enterprise-grade School ERP SaaS platform targeting 1000+ schools. This repository contains a Go backend designed for production deployment on AWS EKS with PostgreSQL, Redis, and observability.

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
3. Set required environment variables:
   ```bash
   export DATABASE_DSN="postgres://user:pass@localhost:5432/school_erp?sslmode=disable"
   export JWT_SECRET="your-secret"
   ```
4. Run local server:
   ```bash
   go run ./cmd/server/main.go
   ```

## Docker

Build the container locally:

```bash
docker build -t school-awesome:latest .
```

Run locally:

```bash
docker run --rm -p 8080:8080 \
  -e DATABASE_DSN="postgres://user:pass@localhost:5432/school_erp?sslmode=disable" \
  -e JWT_SECRET="your-secret" \
  school-awesome:latest
```

## GitHub Actions

The repository includes a GitHub Actions workflow at `.github/workflows/ci.yml` that:

- checks out the repository
- sets up Go 1.24
- caches Go modules
- installs dependencies
- runs `go test ./...`
- builds the Docker image

To publish Docker images from CI, add `DOCKERHUB_USERNAME` and `DOCKERHUB_TOKEN` secrets.

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
