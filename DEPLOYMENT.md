# Deployment Guide

This file explains how to host School Awesome as a web application today.
It covers the backend, frontend, database, and recommended hosting options.

## What you need now

### 1. Backend

The Go backend is containerized via `Dockerfile`.
It reads configuration from `configs/config.yaml` and environment variables.

Required environment variables:
- `DATABASE_DSN` — PostgreSQL connection string
- `JWT_SECRET` — JWT signing secret
- `REDIS_ADDR` — optional Redis address if you use Redis features
- `REDIS_PASSWORD` — optional Redis password
- `OTEL_EXPORTER_OTLP_ENDPOINT` — optional telemetry endpoint

The backend listens on port `8080` by default.

### 2. Frontend

The React frontend is in `frontend/` and is currently served during development by Vite.
For production, build static assets:

```bash
cd frontend
npm install
npm run build
```

Then host the generated files from `frontend/dist` on any static website host:
- Vercel
- Netlify
- GitHub Pages
- AWS S3 + CloudFront
- DigitalOcean App Platform static site

### 3. Database

Use a managed PostgreSQL instance or a hosted PostgreSQL service.
The Docker Compose setup is only for local development.

PostgreSQL must be accessible from the backend with a valid `DATABASE_DSN`.

## Deployment options

### Option A: Full container deployment

1. Build the backend image:

```bash
docker build -t school-awesome-backend:latest .
```

2. Push it to a container registry (Docker Hub, GitHub Container Registry, etc.).
3. Deploy the container on a host that supports Docker:
   - VPS / cloud VM
   - Render
   - Fly.io
   - Railway
   - AWS ECS / Fargate
   - Azure App Service

4. Configure the host with environment variables.
5. Ensure the host can connect to a PostgreSQL database.
6. Expose the backend over HTTPS.

### Option B: Backend container + static frontend

1. Build and deploy the backend container as above.
2. Build the frontend static assets.
3. Deploy the frontend to a static hosting service.
4. Configure the frontend to use the backend URL in production.

Example production API endpoint:
- `https://api.yourschoolapp.com`

Example frontend URL:
- `https://app.yourschoolapp.com`

### Option C: Host backend as binary

1. Build the Go binary locally or on the server.

```bash
GOOS=linux GOARCH=amd64 go build -o server ./cmd/server/main.go
```

2. Transfer `server` and `configs/config.yaml` to the host.
3. Run the backend with environment variables:

```bash
export DATABASE_DSN="postgres://..."
export JWT_SECRET="your-secret"
./server
```

4. Use a reverse proxy like Nginx or Traefik to handle HTTPS and routing.

### Option D: Free self-host on your own machine with a free domain

You can serve the app from your home machine using Docker and expose it publicly with a free domain or tunnel.

#### Local setup

1. Install Docker and Docker Compose on your machine.
2. Run the app stack locally:

```bash
docker compose up --build -d
```

3. Make sure your backend is reachable on `http://localhost:8080` and frontend on `http://localhost:3000`.

#### Free domain options

- **DuckDNS** — free dynamic DNS subdomains like `yourname.duckdns.org`
- **No-IP** — free hostname forwarding with dynamic DNS
- **Freenom** — free domains such as `.tk`, `.ml`, `.ga`, `.cf`, `.gq`
- **Cloudflare** — free DNS if you already have a domain

#### Expose your service publicly

Option 1: Port forwarding
- Forward ports `80` and `443` on your router to your machine
- Use a reverse proxy like Caddy or Nginx to serve the frontend and backend
- Use a free domain from DuckDNS/Freenom and point it to your home IP

Option 2: Cloudflare Tunnel (recommended for home use)
- Install Cloudflare Tunnel on your machine
- Create a tunnel to your local backend and frontend
- Route traffic through a free Cloudflare-managed hostname or your own domain

Option 3: ngrok / localtunnel
- Use `ngrok` or `localtunnel` to expose your local service temporarily
- Good for testing, but free URLs can change frequently

#### Example using Caddy + Docker

1. Use Caddy for HTTPS and reverse proxy.
2. Configure Caddy to forward requests to the local app containers.
3. Point your free domain or DuckDNS hostname to your public IP.

#### Important notes

- Home hosting is okay for testing and demos, but not ideal for production.
- Your ISP may block inbound ports or change your IP frequently.
- Use HTTPS whenever possible.
- Keep secrets like `JWT_SECRET` out of source control.

## What to change for production

- Set `NODE_ENV=production` for frontend build.
- Use a secure `JWT_SECRET` and never commit it.
- Use a production PostgreSQL database, not local Docker Postgres.
- Add TLS/HTTPS for both frontend and backend.
- Use a dedicated domain and DNS.
- Enable backups for the database.

## Recommended hosting route today

For the fastest deployment:
- Deploy backend as a container on a host that supports Docker.
- Deploy frontend as static files on Vercel/Netlify.
- Use a managed PostgreSQL service.
- Add HTTPS via the host provider or a reverse proxy.

## Example production environment

Backend env:
- `DATABASE_DSN=postgres://user:pass@db-host:5432/school_awesome_prod?sslmode=require`
- `JWT_SECRET=super-secret-value`
- `NODE_ENV=production`

Frontend env on build:
- `VITE_BACKEND_URL=https://api.yourschoolapp.com`

## Next steps after hosting

- Configure CORS if the frontend and backend are on different domains.
- Add logging and monitoring.
- Add health endpoint checks.
- Consider a load balancer if needed.
- Plan database migration and seed strategy for production.
