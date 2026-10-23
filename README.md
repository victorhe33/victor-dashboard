# victor-dashboard

A locally run personal dashboard. Go API backend, Vue 3 frontend, wired together with Docker Compose. Right now it serves placeholder widgets; the plumbing is in place for real data sources.

## Stack

| Part | Tech | Path |
|------|------|------|
| API | Go (stdlib `net/http`) | `backend/` |
| UI | Vue 3 + Vite, served by nginx in Docker | `frontend/` |

## Quick start

```sh
docker compose up --build
```

- UI: http://localhost:5173
- API: http://localhost:8080/api/health

## Local development (no Docker)

```sh
# terminal 1
cd backend && go run .

# terminal 2
cd frontend && npm install && npm run dev
```

Vite proxies `/api` to `http://localhost:8080` (set `API_URL` to change it).

## API

| Route | Returns |
|-------|---------|
| `GET /api/health` | `{ "status": "ok", "time": "..." }` |
| `GET /api/widgets` | Placeholder widget list |

## Tests

```sh
cd backend && go test ./...
```
