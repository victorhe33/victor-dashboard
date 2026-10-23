# victor-dashboard

Personal dashboard: Go API (`backend/`) + Vue 3/Vite SPA (`frontend/`). Everything is placeholder scaffolding for now.

## Commands

- Full stack: `make` / `make all` from the repo root (`make server` or `make client` for one service; all wrap `docker compose up --build -d`) (UI and proxied API on `127.0.0.1:3000`; backend has no host port). Services use `restart: unless-stopped`, so this stack is the always-on copy — rebuild it after changes.
- Backend: `cd backend && go run .` / `go test ./...` / `go vet ./...`
- Frontend: `cd frontend && npm run dev` (Vite proxies `/api` to `localhost:8080`; override with `API_URL`) / `npm run build`

## Conventions

- Backend is stdlib-only (`net/http` method routing). Add a dependency only when it earns its place.
- All API routes live under `/api/`; nginx (Docker) and Vite (local) both proxy that prefix to the backend, so keep the frontend on relative `/api/...` URLs.
- Placeholder data lives in `backend/main.go` handlers; replace it there when wiring real sources.
- Use npm (not Yarn) in `frontend/`; keep `package-lock.json` committed since the Docker build runs `npm ci`.
