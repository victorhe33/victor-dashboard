# Rebuild and (re)start the always-on Docker stack.
.PHONY: all server client

all:
	docker compose up --build -d

server:
	docker compose up --build -d backend

client:
	docker compose up --build -d frontend
