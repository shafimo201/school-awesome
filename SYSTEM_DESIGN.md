# School Awesome System Design

## Overview

School Awesome is a local School ERP SaaS prototype with a Go backend, a React/Vite frontend, and PostgreSQL persistence. The system supports authentication, role-based access, user management, and admin workflows for creating students and teachers.

## Architecture Diagram

```text
+---------------------+        +-----------------------+      +-------------------+
|                     |        |                       |      |                   |
|  React Frontend     | <----> |  Go Backend (Gin)     | <--> |  PostgreSQL DB    |
|  (localhost:3000)   |        |  (localhost:8080)     |      |  (db-data volume) |
|                     |        |                       |      |                   |
+---------------------+        +-----------------------+      +-------------------+
        ^  |                            ^  |                        ^
        |  |                        Auth |                        Data
        |  |                            |  |                        |
        |  +-- HTTP API requests -------+  +--- SQL queries --------+
        |                              
        +-- UI routing, login, admin pages
```

The frontend runs in development mode and proxies API requests to the backend. The backend enforces authentication with JWTs and stores user data in PostgreSQL.

## Core Components

### Backend Service (`app`)
- Language: Go
- Web framework: Gin
- Responsibilities:
  - Authentication and JWT issuance
  - User registration and login
  - Admin user management
  - Health checks
  - Request logging
- Key concepts:
  - Clean architecture with domain, use case, adapter, and infrastructure layers
  - `internal/core/domain` for user models
  - `internal/core/usecase` for business logic
  - `internal/adapter/api` for HTTP request handling
  - `internal/adapter/db` for database persistence

### Frontend Application (`frontend`)
- Framework: React + Vite
- UI library: Material UI (MUI)
- Responsibilities:
  - Login screen
  - Admin dashboard and user forms
  - Routing and route guards
  - Fetching profile and admin APIs
- Notes:
  - Frontend runs on `http://localhost:3000`
  - It proxies requests to the backend service during development

### Database (`db`)
- Engine: PostgreSQL 15
- Purpose:
  - Persist users and audit fields
  - Store passwords as bcrypt hashes
  - Support query by school and user ID
- Local persistence: Docker volume `db-data`

## Authentication Flow

1. User submits login form with `username` and `password`.
2. Frontend sends `POST /api/v1/auth/login`.
3. Backend validates credentials via `UserService.Authenticate`.
4. If valid, backend issues a JWT token and expiration time.
5. Frontend stores the token and uses it for protected requests.
6. Protected endpoints validate the token via `AuthMiddleware`.

## Authorization Model

- Role-based access control is enforced in the backend.
- Admin-only actions are protected by `AdminMiddleware`.
- Admin endpoints:
  - `POST /api/v1/admin/students`
  - `POST /api/v1/admin/teachers`
- Normal authenticated users can access:
  - `GET /api/v1/me`
  - `GET /api/v1/users`

## API Surface

### Public endpoints
- `GET /health` — service health
- `POST /api/v1/auth/login` — login with username/password

### Authenticated endpoints
- `GET /api/v1/me` — current profile
- `GET /api/v1/users` — list school users

### Admin endpoints
- `POST /api/v1/admin/students` — create a student
- `POST /api/v1/admin/teachers` — create a teacher

## Data Model

### User
- `id` — unique identifier
- `email` — used as `username` in the current system
- `full_name` — display name
- `password_hash` — bcrypt hash
- `role_id` — `admin`, `teacher`, or `student`
- `status` — active/suspended
- `last_login_at` — optional timestamp
- `school_id` — multi-tenant support

## Deployment Architecture

### Local Docker Compose
- `db` container runs PostgreSQL
- `app` container runs the Go backend on port `8080`
- `frontend` container runs the React app on port `3000`
- Backend environment variables:
  - `DATABASE_DSN`
  - `JWT_SECRET`

### Access points
- UI: `http://localhost:3000`
- Backend API: `http://localhost:8080`

## Seeded Local Admin

The system seeds a default local admin at startup if it does not exist:
- Username: `admin`
- Password: `Shafi@123`

This supports quick local testing and admin access without manual setup.

## System Design Considerations

### Clean architecture
- Separates business rules from transport and persistence.
- Makes user service logic independent of HTTP and database details.

### Local dev friendliness
- Docker Compose exposes all services on localhost
- React frontend hot reloads with local code changes
- Database persists in a named volume between restarts

### Extensibility
- Additional entities (classes, grades, attendance) can be added as domain models
- More backend APIs can be registered under `internal/adapter/api`
- Frontend pages can be added using route guards and existing auth state

## Future improvements

- Rename `email` field to `username` in domain and DB schema for clarity
- Add refresh token support
- Add user role and permission management pages
- Implement proper error payloads and validation feedback on the frontend
- Add automated health checks and readiness probes for production
- Introduce database migrations with version tracking and seed scripts
