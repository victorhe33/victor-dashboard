# victor-dashboard

Personal dashboard: Go API (`backend/`) + Vue 3/Vite SPA (`frontend/`). Everything is placeholder scaffolding for now.

## Commands

Run `make help` for the full target list; prefer `make` targets over raw commands.

- Docker stack: `make up` (or `make up-backend` / `make up-frontend`), `make down`, `make logs`. UI and proxied API on `127.0.0.1:4580`; backend has no host port. Services use `restart: unless-stopped`, so this stack is the always-on copy — rebuild it after changes.
- Backend: `make test` (vet + test), `make dev-backend` (`go run .` on :4581)
- Frontend: `make dev-frontend` (Vite on :4582, proxies `/api` to `localhost:4581`; override with `API_URL`) / `cd frontend && npm run build`

## Conventions

- Backend is stdlib-only (`net/http` method routing). Add a dependency only when it earns its place.
- All API routes live under `/api/`; nginx (Docker) and Vite (local) both proxy that prefix to the backend, so keep the frontend on relative `/api/...` URLs.
- Placeholder data lives in `backend/main.go` handlers; replace it there when wiring real sources.
- Use npm (not Yarn) in `frontend/`; keep `package-lock.json` committed since the Docker build runs `npm ci`.
