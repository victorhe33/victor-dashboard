.DEFAULT_GOAL := help
.PHONY: help up up-backend up-frontend down logs test dev-backend dev-frontend

help: ## List targets
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "  %-14s %s\n", $$1, $$2}'

# Docker stack (always-on copy at http://localhost:3000)

up: ## Rebuild and start both services
	docker compose up --build -d

up-backend: ## Rebuild and restart the backend container
	docker compose up --build -d backend

up-frontend: ## Rebuild and restart the frontend container
	docker compose up --build -d frontend

down: ## Stop the stack
	docker compose down

logs: ## Follow container logs
	docker compose logs -f

# Local development (no Docker)

test: ## Vet and test the backend
	cd backend && go vet ./... && go test ./...

dev-backend: ## Run the Go API on :8080
	cd backend && go run .

dev-frontend: ## Run the Vite dev server on :5173
	cd frontend && npm run dev
