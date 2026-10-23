# victor-dashboard

A locally run personal dashboard. Go API backend, Vue 3 frontend, wired together with Docker Compose. Right now it serves placeholder widgets; the plumbing is in place for real data sources.

## Stack

| Part | Tech | Path |
|------|------|------|
| API | Go (stdlib `net/http`) | `backend/` |
| UI | Vue 3 + Vite, served by nginx in Docker | `frontend/` |

## Quick start

```sh
make up       # rebuild and start both services
make help     # list all targets
```

- UI: http://localhost:4580
- API (via the nginx proxy): http://localhost:4580/api/health

Ports bind to `127.0.0.1` only, so the dashboard isn't reachable from your network.

## Make targets

| Target | Does |
|--------|------|
| `make up` | Rebuild and start both containers |
| `make up-backend` / `make up-frontend` | Rebuild and restart one container |
| `make down` | Stop the stack |
| `make logs` | Follow container logs |
| `make test` | `go vet` + `go test` the backend |
| `make dev-backend` / `make dev-frontend` | Run locally without Docker |

Plain `make` prints the target list. Docker targets wrap `docker compose`, so any compose command still works directly.

### Run on login

Both services use `restart: unless-stopped`, so once started they come back whenever Docker does. Enable Docker Desktop → Settings → General → **Start Docker Desktop when you sign in** and the dashboard will be up after every reboot. `make down` stops it until the next `make up`.

To pick up code changes, rerun `make up` (or `make up-backend` / `make up-frontend`).

## Local development (no Docker)

```sh
cd frontend && npm install   # first time only

make dev-backend    # terminal 1
make dev-frontend   # terminal 2
```

UI at http://localhost:4582, API at http://localhost:4581. Vite proxies `/api` to `http://localhost:4581` (set `API_URL` to change it). These ports don't overlap with the Docker stack, so both can run at once.

## API

| Route | Returns |
|-------|---------|
| `GET /api/health` | `{ "status": "ok", "time": "..." }` |
| `GET /api/widgets` | Placeholder widget list |

## Tests

```sh
make test
```
