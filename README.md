# victor-dashboard

A locally run personal dashboard. Go API backend, Vue 3 frontend, wired together with Docker Compose. Right now it serves placeholder widgets; the plumbing is in place for real data sources.

## Stack

| Part | Tech | Path |
|------|------|------|
| API | Go (stdlib `net/http`) | `backend/` |
| UI | Vue 3 + Vite, served by nginx in Docker | `frontend/` |

## Quick start

```sh
docker compose up --build -d
```

- UI: http://localhost:3000
- API (via the nginx proxy): http://localhost:3000/api/health

Ports bind to `127.0.0.1` only, so the dashboard isn't reachable from your network.

### Run on login

Both services use `restart: unless-stopped`, so once started they come back whenever Docker does. Enable Docker Desktop → Settings → General → **Start Docker Desktop when you sign in** and the dashboard will be up after every reboot. `docker compose down` stops it until the next `up`.

To pick up code changes, rerun `docker compose up --build -d`.

## Local development (no Docker)

```sh
# terminal 1
cd backend && go run .

# terminal 2
cd frontend && npm install && npm run dev
```

UI at http://localhost:5173, API at http://localhost:8080. Vite proxies `/api` to `http://localhost:8080` (set `API_URL` to change it). These ports don't overlap with the Docker stack, so both can run at once.

## API

| Route | Returns |
|-------|---------|
| `GET /api/health` | `{ "status": "ok", "time": "..." }` |
| `GET /api/widgets` | Placeholder widget list |

## Tests

```sh
cd backend && go test ./...
```
