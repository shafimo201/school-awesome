#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${DATABASE_DSN:-}" ]]; then
  echo "DATABASE_DSN is required"
  exit 1
fi

if [[ -z "${JWT_SECRET:-}" ]]; then
  echo "JWT_SECRET is required"
  exit 1
fi

go run ./cmd/server/main.go
