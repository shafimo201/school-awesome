# Free Hosting Plan

This guide explains how to host School Awesome for free using available services and tools.
It focuses on the easiest path for a web app and mobile-ready deployment.

## Goal

Host the backend, frontend, and mobile-friendly access with minimal cost.

## Recommended free hosting approach

### 1. Frontend hosting (free)

Use Vercel or Netlify for the React app.
They both offer a free tier with automatic deployment from GitHub.

- Vercel: deploy `frontend/` as a static site
- Netlify: deploy `frontend/dist` after building

### 2. Backend hosting (free tier)

Use a free container or serverless platform:
- **Railway** free tier
- **Render** free tier for web services
- **Fly.io** free tier with small VM
- **Heroku** free tier (legacy, may not be ideal)

These can deploy your backend container or direct from source.

### 3. Database hosting (free or low-cost)

Use a free or trial-tier PostgreSQL service:
- **Railway** free PostgreSQL add-on
- **Supabase** free Postgres database
- **Neon** free Postgres
- **ElephantSQL** free plan

If not available, you can also host a PostgreSQL container on the same free platform if supported.

### 4. Free domain options

- **Freenom** — free domains like `.tk`, `.ml`, `.ga`, `.cf`, `.gq`
- **DuckDNS** — free dynamic DNS hostnames
- **No-IP** — free hostname forwarding
- **Cloudflare** — free DNS service if you have a domain

### Free hosting workflow

#### Step 1: Deploy the backend

1. Build or connect your repo to the platform.
2. Set environment variables:
   - `DATABASE_DSN`
   - `JWT_SECRET`
3. Ensure backend is reachable as a public URL.

Example (Render / Railway):
- `https://school-awesome-backend.onrender.com`

#### Step 2: Deploy the frontend

1. Connect the `frontend/` folder to Vercel or Netlify.
2. Configure build command:
   - `npm install && npm run build`
3. Set the production backend URL as an environment variable.

Example:
- `VITE_BACKEND_URL=https://school-awesome-backend.onrender.com`

#### Step 3: Use a free domain

1. Register a free domain on Freenom or get a DuckDNS hostname.
2. Point the domain to your frontend site.
3. For backend, either use the provider URL or create a subdomain / CNAME.

### Free hosting notes

- Vercel/Netlify are easiest for frontend and offer HTTPS automatically.
- Railway/Render/Fly.io are easiest for backend deployment.
- Supabase/Neon are easiest for free Postgres data.
- For a home-hosted backend, use Cloudflare Tunnel with a free hostname.

## Free mobile access

### Option A: Expo Go
- Build a mobile client with Expo.
- Use Expo Go to run the mobile app for free.
- No store publishing required for testing.

### Option B: PWA
- Make the existing web app mobile-friendly.
- Host the PWA on Vercel or Netlify for free.

## Example free stack

- Frontend: `https://your-school-app.vercel.app`
- Backend: `https://school-awesome-backend.up.railway.app`
- Database: free Postgres on Supabase
- Domain: `https://your-school-app.tk`

## Checklist

- [ ] Backend deployed on free tier service
- [ ] Frontend deployed on free static host
- [ ] Database running on free Postgres service
- [ ] Free domain registered and configured
- [ ] Frontend configured to call backend URL
- [ ] HTTPS enabled on frontend and backend

## Recommended path for you

1. Use **Vercel** for frontend.
2. Use **Railway** or **Render** for backend.
3. Use **Supabase** or **Neon** for PostgreSQL.
4. If you need a custom free domain, use **Freenom**.
5. If using home hosting instead, use **Cloudflare Tunnel**.

## Helpful links

- Vercel: https://vercel.com
- Netlify: https://netlify.com
- Railway: https://railway.app
- Render: https://render.com
- Fly.io: https://fly.io
- Supabase: https://supabase.com
- Neon: https://neon.tech
- Freenom: https://freenom.com
- DuckDNS: https://duckdns.org
- Cloudflare Tunnel: https://developers.cloudflare.com/cloudflare-one/connections/connect-apps
